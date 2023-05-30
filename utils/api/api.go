package api

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Api struct {
	server  *gin.Engine
	dbConn  *gorm.DB
	routers []Router
}

type Router interface {
	Route(handler *gin.RouterGroup)
}

func (a Api) Start() error {
	root := a.server.Group("/main")
	for _, router := range a.routers {
		router.Route(root)
	}

	var err error
	if err = a.server.Run(); err != nil {
		return err
	}

	return nil
}
