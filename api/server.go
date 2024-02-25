package main

import (
	"github.com/4rgetlahm/backports/api/database"
	"github.com/gin-gonic/gin"
)

func Init() {
	database.Init()
}

func main() {
	Init()
	middleware.InitAuthenticationMiddleware()

	server := gin.Default()

	server.Use(gin.Logger())
	server.Use(gin.Recovery())
	server.Use(middleware.CORS())
	controller.BindTokenController(server)
	controller.BindEventRoutes(server)

	server.Run(":5000")
}
