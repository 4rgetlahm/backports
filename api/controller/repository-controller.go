package controller

import (
	"github.com/4rgetlahm/backports/api/service"
	"github.com/gin-gonic/gin"
)

func InitRepositoryController(router *gin.Engine) {
	router.GET("/v1/repositories", GetRepositories)
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
	name := c.Param("name")

	repository, err := service.GetRepository(name)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, repository)
}

func CreateRepository(c *gin.Context) {
	type CreateRepositoryRequest struct {
		VersionControlSystem string `json:"version_control_system"`
		CloneURL             string `json:"clone_url"`
		Name                 string `json:"name"`
	}

	var request CreateRepositoryRequest
	err := c.BindJSON(&request)

	if err != nil {
		c.JSON(400, gin.H{"error": "invalid request"})
		return
	}

	repository, err := service.CreateRepository(request.VersionControlSystem, request.CloneURL, request.Name)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, repository)
}
