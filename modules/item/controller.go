package item

import (
	"context"
	"fmt"
	"github.com/dibimbing-satkom-indo/gilded-roses-store-go/dto"
	"time"
)

type Controller struct {
	usecase UsecaseInterface
}

type ControllerInterface interface {
	CreateItem(
		ctx context.Context,
		payload CreateItemPayload,
	) (*dto.Response, error)
}

func (c Controller) CreateItem(
	ctx context.Context,
	payload CreateItemPayload,
) (*dto.Response, error) {
	start := time.Now()
	item, err := c.usecase.CreateItem(ctx, payload)
	if err != nil {
		return nil, err
	}

	return dto.NewSuccessResponse(
		item,
		"item created successfully",
		fmt.Sprint(time.Since(start).Milliseconds(), ".ms"),
	), nil
}
