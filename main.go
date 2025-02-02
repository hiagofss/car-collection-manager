package main

import (
	routes "car-collection-manager/router"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	routes.SetupRouter(r)
}
