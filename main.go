package main

import (
	"github.com/rrojan/gin-skeleton/api"
)

func main() {
	app := api.NewApp()
	app.Run()
}
