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
	"io/ioutil"
	"log"
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
func CreateKerbal(ctx context.Context, items KerbalItems) error {
	id := generateID(items)
	images, err := loadImages(ctx, items)
	if err != nil {
		return nil
	}
	var kerbalBuf bytes.Buffer
	if err := drawImage(ctx, images, &kerbalBuf); err != nil {
		return err
	}
	if err := ioutil.WriteFile("./kerbal.png", kerbalBuf.Bytes(), 0644); err != nil {
		return err
	}
	fmt.Println(id)
	return nil
}

func loadImages(ctx context.Context, items KerbalItems) ([]image.Image, error) {
	awsService := services.New(ctx)
	if err := awsService.AWSConnect(); err != nil {
		return nil, errors.Wrap(err, "Failed to connect to aws: %s")
	}
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
			img, _, err := awsService.DownloadImages(fmt.Sprintf("/%s/%s", folder, items[item]))
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
	for folder, item := range items {
		if item == "" {
			continue
		}
		name += fmt.Sprintf("%s:%s", folder, item)
	}
	bytes := hash.Sum(nil)
	return fmt.Sprintf("%x\n", bytes)
}
