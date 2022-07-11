package main

import (
	"libary/model"
	"libary/routes"

	"github.com/gin-gonic/gin"
)

func init() {
	model.ConnectDatabase()
}

func main() {

	router := gin.Default()

	routes.Groups(router)

	router.Run("localhost:8080")

}
