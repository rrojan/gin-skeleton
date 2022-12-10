package main

import (
	"github.com/rrojan/gin-skeleton/api/routes"
	"github.com/rrojan/gin-skeleton/db"
)

func main() {
	db.Connect()
	r := routes.NewRouter()
	r.Run(":8000")
}
