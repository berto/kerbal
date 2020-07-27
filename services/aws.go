package services

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/dustin/go-humanize"
	"github.com/hako/durafmt"
	"github.com/pkg/errors"
)

// Service for aws bucket
type Service struct {
	ctx    context.Context
	Sess   *session.Session
	Bucket string
	Prefix string
}

// S3Object holds important information from an s3 key
type S3Object struct {
	svc        *Service
	Bucket     string    `json:"bucket"`
	Name       string    `json:"name"`
	Date       time.Time `json:"date"`
	PrettyDate string    `json:"pretty_date"`
	Size       int64     `json:"size"`
	Tag        string    `json:"tag"`
	Age        string    `json:"age"`
}

func (s S3Object) String() string {
	ju, _ := json.MarshalIndent(s, " ", "")
	return string(ju)
}

// New aws services
func New(ctx context.Context) *Service {
	return &Service{ctx: ctx}
}

// AWSConnect establishes aws connection
func (s *Service) AWSConnect() error {
	profile := os.Getenv("AWS_PROFILE")
	if profile == "" {
		profile = "kerbal.me"
	}
	conf := aws.NewConfig().WithCredentials(
		credentials.NewSharedCredentials("", profile),
	).WithRegion("us-west-2")
	sess, err := session.NewSession(conf)
	if err != nil {
		return err
	}
	s.Sess = sess
	s.Bucket = profile
	return nil
}

// DownloadBytes downloads a key into bytes
func (s *Service) DownloadBytes(keyName string) ([]byte, error) {
	downloader := s3manager.NewDownloader(s.Sess)
	buffer := &aws.WriteAtBuffer{}
	path := strings.Join([]string{s.Prefix, keyName}, "/")
	_, err := downloader.Download(buffer, &s3.GetObjectInput{
		Bucket: aws.String(s.Bucket),
		Key:    aws.String(path),
	})
	if err != nil {
		return nil, errors.Wrapf(
			err,
			"download manager: bucket: %s prefix: %s keyname: %s",
			s.Bucket,
			s.Prefix,
			keyName,
		)
	}
	data, err := ioutil.ReadAll(bytes.NewReader(buffer.Bytes()))
	if err != nil {
		return nil, errors.Wrap(err, "reading aws buffer")
	}
	return data, nil
}

// List lists the s3 objects at a bucket and prefix
func (s *Service) List() ([]*S3Object, error) {
	client := s3.New(s.Sess)
	data, err := client.ListObjects(&s3.ListObjectsInput{
		Bucket: aws.String(s.Bucket),
	})
	if err != nil {
		return nil, errors.Wrap(err, "listing s3 objects")
	}
	out := []*S3Object{}
	for _, obj := range data.Contents {
		out = append(out, s.NewS3Object(obj))
	}
	return out, nil
}

// NewS3Object creates a new s3 object
func (s *Service) NewS3Object(obj *s3.Object) *S3Object {
	return &S3Object{
		svc:        s,
		Bucket:     s.Bucket,
		Name:       *obj.Key,
		Date:       *obj.LastModified,
		PrettyDate: humanize.Time(*obj.LastModified),
		Size:       *obj.Size,
		Tag:        *obj.ETag,
		Age:        fmt.Sprint(durafmt.ParseShort(time.Since(*obj.LastModified))),
	}
}

// UploadFromReader uploads to bucket from io reader
func (s *S3Object) UploadFromReader(reader io.Reader) error {
	uploader := s3manager.NewUploader(s.svc.Sess)
	_, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(s.Bucket),
		Key:    aws.String(s.Name),
		Body:   reader,
	})
	if err != nil {
		return errors.Wrap(err, "upload")
	}
	return nil
}
