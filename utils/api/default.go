package api

import (
	"fmt"
	"github.com/dibimbing-satkom-indo/gilded-roses-store-go/modules/health"
	"github.com/dibimbing-satkom-indo/gilded-roses-store-go/modules/item"
	"github.com/dibimbing-satkom-indo/gilded-roses-store-go/utils/buckets"

	"github.com/dibimbing-satkom-indo/gilded-roses-store-go/utils/db"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

func Default() *Api {
	dbConn, err := db.Default()
	if err != nil {
		panic(fmt.Errorf("db.Default: %w", err))
	}

	bucket := buckets.Default()
	routers := []Router{
		health.NewRouter(dbConn),
		item.NewRouter(dbConn, bucket),
	}

	server := gin.Default()
	return &Api{
		server:  server,
		routers: routers,
	}
}
