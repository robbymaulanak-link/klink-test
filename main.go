package main

import (
	"test-k-link-indonesia/config"
	"test-k-link-indonesia/packages/connection"
	"test-k-link-indonesia/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	connection.Database()

	config.Migration()

	routes.RouteInit(r.Group("k-link/api/v1"))

	r.Run(":5000")
}
