package controller

import (
	"github.com/4rgetlahm/backports/api/service"
	"github.com/gin-gonic/gin"
)

func InitRepositoryController(router *gin.Engine) {
	router.GET("/v1/repositories", GetRepositories)
	router.GET("/v1/repository/:owner/:name", GetRepository)
	router.POST("/v1/repository", CreateRepository)
}

func GetRepositories(c *gin.Context) {
	repositories, err := service.GetRepositories()

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, repositories)
}

func GetRepository(c *gin.Context) {
	owner := c.Param("owner")
	name := c.Param("name")

	repository, err := service.GetRepository(owner, name)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, repository)
}

func CreateRepository(c *gin.Context) {
	type CreateRepositoryRequest struct {
		CloneURL string `json:"clone_url"`
		Image    string `json:"image"`
	}

	var request CreateRepositoryRequest
	err := c.BindJSON(&request)

	if err != nil {
		c.JSON(400, gin.H{"error": "invalid request"})
		return
	}

	repository, err := service.CreateRepository(request.CloneURL)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, repository)
}
