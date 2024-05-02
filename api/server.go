package main

import (
	"github.com/4rgetlahm/backports/api/controller"
	"github.com/4rgetlahm/backports/api/localGRPC"
	"github.com/4rgetlahm/backports/api/middleware"
	"github.com/4rgetlahm/backports/database"
	"github.com/gin-gonic/gin"
)

func main() {
	database.Init()

	conn := localGRPC.Init()
	defer conn.Close()

	server := gin.Default()

	controller.InitBackportController(server)
	controller.InitRepositoryController(server)

	server.Use(gin.Logger())
	server.Use(gin.Recovery())
	server.Use(middleware.CORS())

	server.Run(":5000")
}
