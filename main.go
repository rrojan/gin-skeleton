package main

import (
	"github.com/rrojan/gin-skeleton/api/routes"
)

func main() {
	r := routes.NewRouter()
	r.Run(":8000")
}
