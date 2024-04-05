package controllers

import (
	"awesomeProject1/initializers"
	"awesomeProject1/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func BillboardCreate(c *gin.Context) {
	// Get data off req body
	var body struct {
		Name        string
		Description string
		State       string
		Status      string
	}

	if c.BindJSON(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}
	// Create a Client
	billboard := models.Billboard{Name: body.Name, Description: body.Description, Status: body.Status, State: body.State}

	result := initializers.DB.Create(&billboard)

	if result.Error != nil {
		c.Status(400)
		return
	}
	// Return it

	c.JSON(200, gin.H{
		"billboard": billboard,
	})
}

func BillboardIndex(c *gin.Context) {

	//Get the posts
	var billboards []models.Billboard
	initializers.DB.Find(&billboards)

	// Response with them
	c.JSON(200, gin.H{
		"billboards": billboards,
	})
}

func BillboardShow(c *gin.Context) {
	//Get id off url
	id := c.Param("id")

	//Get the posts
	var billboard models.Billboard
	initializers.DB.First(&billboard, id)

	// Response with them
	c.JSON(200, gin.H{
		"billboard": billboard,
	})
}

func BillboardUpdate(c *gin.Context) {
	//Get the id off the url
	Id := c.Param("Id")

	//Get the data off req body
	var body struct {
		Name        string
		Description string
		State       string
		Status      string
	}

	c.Bind(&body)

	//Find the post were upadating
	var billboard models.Billboard
	initializers.DB.First(&billboard, Id)

	//Updating it
	initializers.DB.Model(&billboard).Updates(models.Billboard{
		Name:        body.Name,
		Description: body.Description,
		Status:      body.Status,
		State:       body.State,
	})

	//Response with it
	c.JSON(200, gin.H{
		"billboard": billboard,
	})
}

func BillboardDelete(c *gin.Context) {
	//Get the id off the url
	Id := c.Param("Id")

	//Delete the posts
	initializers.DB.Delete(&models.Billboard{}, Id)

	//Response
	c.Status(200)
}
