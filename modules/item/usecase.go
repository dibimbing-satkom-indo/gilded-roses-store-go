package item

import (
	"context"
	"fmt"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/dibimbing-satkom-indo/gilded-roses-store-go/entities"
	"github.com/dibimbing-satkom-indo/gilded-roses-store-go/repositories"
	"github.com/dibimbing-satkom-indo/gilded-roses-store-go/utils/buckets"
	"mime/multipart"
	"sync"
	"time"
)

type Usecase struct {
	bucket    buckets.Uploader
	itemRepo  repositories.ItemRepositoryInterface
	imageRepo repositories.ImageRepositoryInterface
}

type UsecaseInterface interface {
	CreateItem(ctx context.Context, payload CreateItemPayload) (*entities.Items, error)
}

func (u Usecase) CreateItem(ctx context.Context, payload CreateItemPayload) (*entities.Items, error) {
	//TODO: Insert file to table item
	item, err := u.itemRepo.Store(ctx, &entities.Items{
		Name:    payload.Nama,
		SellIn:  payload.SellIn,
		Quality: payload.Quality,
	})
	if err != nil {
		return nil, err
	}

	// TODO: upload images to cloudinary
	var (
		wg         = new(sync.WaitGroup)
		errChan    = make(chan error)
		res        *uploader.UploadResult
		imagesChan = make(chan entities.Images, len(payload.Images))
		images     = make([]entities.Images, len(payload.Images))
	)

	wg.Add(len(payload.Images))
	for i, image := range payload.Images {
		fileName := fmt.Sprintf("item-%s-%d", item.Name, time.Now().Unix())
		go func(image multipart.File, fileName string, i int) {
			defer wg.Done()
			res, err = u.bucket.Upload(ctx, image, uploader.UploadParams{
				PublicID: fileName,
			})
			imagesChan <- entities.Images{
				ItemID: item.ID,
				Url:    res.SecureURL,
			}
			errChan <- err
		}(image, fileName, i)
	}

	for i := 0; i < len(payload.Images); i++ {
		err = <-errChan
		if err != nil {
			return nil, err
		}
	}

	for i := 0; i < len(payload.Images); i++ {
		image := <-imagesChan
		images[i] = image
	}

	wg.Wait()

	// TODO: for the the url, insert to table images
	images, err = u.imageRepo.BatchStore(ctx, images)
	if err != nil {
		return nil, err
	}

	item.Images = images

	return item, nil
}
