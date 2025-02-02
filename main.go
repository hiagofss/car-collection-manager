package main

import (
	"car-collection-manager/database"
	routes "car-collection-manager/router"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	database.ConnectMongoDb()

	routes.SetupRouter(r)
}
