package main

import (
	"fmt"
	"github.com/dibimbing-satkom-indo/gilded-roses-store-go/utils/api"
)

func main() {
	app := api.Default()
	if err := app.Start(); err != nil {
		panic(fmt.Errorf("app.Start: %w", err))
	}
}
