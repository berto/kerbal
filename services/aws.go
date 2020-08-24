package services

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"image"
	"io"
	"os"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
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
	region := os.Getenv("AWS_REGION")
	if region == "" {
		region = "us-west-2"
	}
	bucket := os.Getenv("AWS_BUCKET")
	if bucket == "" {
		bucket = "kerbal.me"
	}
	conf := aws.NewConfig().WithRegion(region)
	sess, err := session.NewSession(conf)
	if err != nil {
		return err
	}
	s.Sess = sess
	s.Bucket = bucket
	return nil
}

// DownloadImages downloads images from bucket
func (s *Service) DownloadImages(keyName string) (image.Image, string, error) {
	downloader := s3manager.NewDownloader(s.Sess)
	buffer := &aws.WriteAtBuffer{}
	path := strings.Join([]string{s.Prefix, keyName}, "/")
	_, err := downloader.Download(buffer, &s3.GetObjectInput{
		Bucket: aws.String(s.Bucket),
		Key:    aws.String(path),
	})
	if err != nil {
		return nil, "", errors.Wrapf(
			err,
			"download manager: bucket: %s prefix: %s keyname: %s",
			s.Bucket,
			s.Prefix,
			keyName,
		)
	}
	return image.Decode(bytes.NewReader(buffer.Bytes()))
}

// List lists the s3 objects at a bucket and prefix
func (s *Service) List(prefix *string) ([]*S3Object, error) {
	client := s3.New(s.Sess)
	data, err := client.ListObjectsV2(&s3.ListObjectsV2Input{
		Bucket: aws.String(s.Bucket),
		Prefix: prefix,
	})
	if err != nil {
		return nil, errors.Wrap(err, "listing s3 objects")
	}
	out := []*S3Object{}
	for _, obj := range data.Contents {
		out = append(out, s.NewS3FromObj(obj))
	}
	return out, nil
}

// NewS3FromObj creates a new s3 object
func (s *Service) NewS3FromObj(obj *s3.Object) *S3Object {
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

// NewS3Object from key
func (s *Service) NewS3Object(keyName string) *S3Object {
	return &S3Object{
		svc:    s,
		Bucket: s.Bucket,
		Name:   keyName,
		Date:   time.Now(),
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
