package item

import (
	"github.com/dibimbing-satkom-indo/gilded-roses-store-go/repositories"
	"github.com/dibimbing-satkom-indo/gilded-roses-store-go/utils/buckets"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Router struct {
	rq RequestHandlerInterface
}

func NewRouter(
	db *gorm.DB,
	bucket buckets.Uploader,
) Router {
	return Router{
		rq: &RequestHandler{
			ctrl: Controller{
				usecase: Usecase{
					bucket:    bucket,
					itemRepo:  repositories.NewItemRepository(db),
					imageRepo: repositories.NewImagesRepository(db),
				},
			},
		},
	}
}

func (r Router) Route(handler *gin.RouterGroup) {
	item := handler.Group("/inventory")
	item.POST(
		"/",
		r.rq.CreateItem,
	)
}
