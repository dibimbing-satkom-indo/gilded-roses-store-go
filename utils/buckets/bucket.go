package buckets

import (
	"context"
	"fmt"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"os"
)

type Bucket struct {
	client *cloudinary.Cloudinary
}

type Uploader interface {
	Upload(
		ctx context.Context,
		file interface{},
		uploadParams uploader.UploadParams,
	) (*uploader.UploadResult, error)
}

func Default() Uploader {
	cloudName := os.Getenv("CLOUDINARY_CLOUD_NAME")
	apiKey := os.Getenv("CLOUDINARY_API_KEY")
	secret := os.Getenv("CLOUDINARY_API_SECRET")

	cld, err := cloudinary.NewFromParams(cloudName, apiKey, secret)
	if err != nil {
		panic(fmt.Sprintf("cloudinary.NewFromParams: %s", err.Error()))
	}
	return &Bucket{
		client: cld,
	}
}

func (cld Bucket) Upload(
	ctx context.Context,
	file interface{},
	uploadParams uploader.UploadParams,
) (*uploader.UploadResult, error) {
	res, err := cld.client.Upload.Upload(ctx, file, uploadParams)
	if err != nil {
		return &uploader.UploadResult{}, err
	}

	return res, nil
}
