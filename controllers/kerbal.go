package controllers

import (
	"bytes"
	"context"
	"crypto/sha1"
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"io"
	"log"
	"strings"
	"sync"

	"github.com/berto/kerbal/services"
	"github.com/pkg/errors"
)

// KerbalItems is a list of avatar items
type KerbalItems map[string]string

// Validate checks for required items
func (k KerbalItems) Validate() error {
	for item, value := range k {
		if requiredItems[item] && value == "" {
			return errors.New(fmt.Sprintf("%s item is required", item))
		}
	}
	return nil
}

var availableItems = []string{
	"suit",
	"color",
	"eyes",
	"mouth",
	"hair",
	"facial-hair",
	"glasses",
	"extras",
	"suit-front",
}

var requiredItems = map[string]bool{
	"color": true,
	"eyes":  true,
	"mouth": true,
	"suit":  true,
}

// CreateKerbal takes a list of items and generates avatar
func CreateKerbal(ctx context.Context, items KerbalItems) (string, error) {
	awsService := services.New(ctx)
	if err := awsService.AWSConnect(); err != nil {
		return "", errors.Wrap(err, "Failed to connect to aws: %s")
	}
	id := generateID(items)
	new, err := isNew(awsService, id)
	if err != nil {
		return "", err
	}
	if !new {
		return id, nil
	}
	images, err := loadImages(awsService, items)
	if err != nil {
		return id, err
	}
	var kerbalBuf bytes.Buffer
	if err := drawImage(ctx, images, &kerbalBuf); err != nil {
		return id, err
	}
	obj := awsService.NewS3Object(fmt.Sprintf("/kerbals/%s.png", id))
	return id, obj.UploadFromReader(bytes.NewReader(kerbalBuf.Bytes()))
}

func isNew(awsService *services.Service, id string) (bool, error) {
	folder := "kerbals"
	kerbalObjs, err := awsService.List(&folder)
	if err != nil {
		return false, errors.Wrap(err, "Failed to list items: %s")
	}
	for _, obj := range kerbalObjs {
		if getName(obj.Name) == id {
			return false, nil
		}
	}
	return true, nil
}

func getName(objName string) string {
	name := strings.Split(objName, ".")[0]
	split := strings.Split(name, "/")
	if len(split) != 2 {
		return ""
	}
	return split[1]
}

func loadImages(awsService *services.Service, items KerbalItems) ([]image.Image, error) {
	images := map[string]image.Image{}
	var wg sync.WaitGroup
	var mtx sync.Mutex
	errs := make(map[string]error)
	for _, item := range availableItems {
		if items[item] == "" {
			continue
		}
		wg.Add(1)
		go func(item string) {
			defer wg.Done()
			folder := item
			if folder == "suit-front" {
				folder = "suit"
			}
			img, _, err := awsService.DownloadImages(fmt.Sprintf("/images/%s/%s", folder, items[item]))
			if err != nil {
				mtx.Lock()
				errs[item] = err
				mtx.Unlock()
				return
			}
			mtx.Lock()
			images[item] = img
			mtx.Unlock()
		}(item)
	}
	wg.Wait()
	if len(errs) > 0 {
		log.Printf("processing images: %s", errs)
		return nil, errors.New("failed to load images")
	}
	imageList := []image.Image{}
	for _, item := range availableItems {
		imageList = append(imageList, images[item])
	}
	return imageList, nil
}

func drawImage(ctx context.Context, images []image.Image, w io.Writer) error {
	var first image.Image
	for _, img := range images {
		if img == nil {
			continue
		}
		first = img
		break
	}
	if first == nil {
		return errors.New("no image found")
	}
	output := image.NewRGBA(first.Bounds())
	for _, img := range images {
		if img == nil {
			continue
		}
		draw.Draw(output, output.Bounds(), img, image.ZP, draw.Over)
	}
	if err := png.Encode(w, output); err != nil {
		return err
	}
	return nil
}

func generateID(items KerbalItems) string {
	hash := sha1.New()
	name := ""
	for _, item := range availableItems {
		if items[item] == "" {
			continue
		}
		name += fmt.Sprintf("%s:%s", item, items[item])
	}
	hash.Write([]byte(name))
	bytes := hash.Sum(nil)
	return fmt.Sprintf("%x", bytes)
}
