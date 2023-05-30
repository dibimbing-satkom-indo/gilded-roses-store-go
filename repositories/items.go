package repositories

import (
	"context"
	"github.com/dibimbing-satkom-indo/gilded-roses-store-go/entities"
	"gorm.io/gorm"
)

type ItemRepository struct {
	db *gorm.DB
}

func NewItemRepository(db *gorm.DB) ItemRepositoryInterface {
	return ItemRepository{db: db}
}

type ItemRepositoryInterface interface {
	Store(ctx context.Context, item *entities.Items) (*entities.Items, error)
}

func (repo ItemRepository) Store(ctx context.Context, item *entities.Items) (*entities.Items, error) {
	err := repo.db.WithContext(ctx).
		Create(&item).
		Error

	return item, err
}
