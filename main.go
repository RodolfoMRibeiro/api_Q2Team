package main

import (
	"libary/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	routes.Groups(router)

	router.Run("localhost:8080")

}
