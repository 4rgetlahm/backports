package controller

import (
	"strconv"

	"github.com/4rgetlahm/backports/api/service"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InitBackportController(router *gin.Engine) {
	router.GET("/v1/backports/:from/:to", GetBackports)
	router.GET("/v1/backport/:id", GetBackport)
	router.POST("/v1/backport", CreateBackport)
}

func GetBackports(c *gin.Context) {
	fromParam := c.Param("from")
	toParam := c.Param("to")

	from, err := strconv.Atoi(fromParam)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid from parameter"})
		return
	}

	to, err := strconv.Atoi(toParam)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid to parameter"})
		return
	}

	backports, err := service.GetBackports(from, to)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, backports)
}

func GetBackport(c *gin.Context) {
	id := c.Param("id")

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid ID"})
		return
	}

	backport, err := service.GetBackport(objectId)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, backport)
}

func CreateBackport(c *gin.Context) {
	type CreateBackportRequest struct {
		Author          string   `json:"author"`
		Commits         []string `json:"commits"`
		RepositoryOwner string   `json:"repositoryOwner"`
		RepositoryName  string   `json:"repositoryName"`
		TargetBranch    string   `json:"targetBranch"`
	}

	var request CreateBackportRequest
	err := c.BindJSON(&request)

	if err != nil {
		c.JSON(400, gin.H{"error": "invalid request"})
		return
	}

	backport, err := service.CreateBackport(request.Author, request.Commits, request.RepositoryOwner, request.RepositoryName, request.TargetBranch)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, backport)
}
