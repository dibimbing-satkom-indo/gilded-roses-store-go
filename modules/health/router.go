package health

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Router struct {
	requestHandler *RequestHandler
}

func NewRouter(db *gorm.DB) *Router {
	return &Router{
		requestHandler: &RequestHandler{repo: versionGetterRepository{db: db}},
	}
}

func (r Router) Route(route *gin.RouterGroup) {
	route.GET("/health", r.requestHandler.check)
}
