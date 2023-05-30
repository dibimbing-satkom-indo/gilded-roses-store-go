package api

import (
	"bitbucket.org/ifan-moladin/base-project/utils/httpserver"
	"gorm.io/gorm"
)

type Api struct {
	httpServer httpserver.HttpServer
	dbConn     *gorm.DB
	routers    []Router
}

type Router interface {
	Route(rgh httpserver.RouteHandler)
}

func (a Api) Start() error {
	root := a.httpServer.Group("/")
	for _, router := range a.routers {
		router.Route(root)
	}

	var err error
	if err = a.httpServer.Run(); err != nil {
		return err
	}

	return nil
}

func New(
	httpServer httpserver.HttpServer,
	dbConn *gorm.DB,
	routers []Router,
) *Api {
	return &Api{
		httpServer: httpServer,
		dbConn:     dbConn,
		routers:    routers,
	}
}
