package repositories

import (
	"context"
	"github.com/dibimbing-satkom-indo/gilded-roses-store-go/entities"
	"gorm.io/gorm"
)

type ImagesRepository struct {
	db *gorm.DB
}

func NewImagesRepository(db *gorm.DB) ImageRepositoryInterface {
	return ImagesRepository{
		db: db,
	}
}

type ImageRepositoryInterface interface {
	BatchStore(
		ctx context.Context,
		images []entities.Images,
	) ([]entities.Images, error)
}

func (repo ImagesRepository) BatchStore(
	ctx context.Context,
	images []entities.Images,
) ([]entities.Images, error) {
	err := repo.db.WithContext(ctx).
		Create(&images).
		Error

	return images, err
}
