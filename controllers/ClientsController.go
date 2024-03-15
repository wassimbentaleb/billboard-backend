package controllers

import (
	"awesomeProject1/initializers"
	"awesomeProject1/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ClientCreate(c *gin.Context) {
	// Get data off req body
	var body struct {
		Email     string
		FirstName string
		LastName  string
		State     string
	}

	if c.BindJSON(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}
	// Create a Client
	post := models.Client{Email: body.Email, FirstName: body.FirstName, LastName: body.LastName, State: body.State}

	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}
	// Return it

	c.JSON(200, gin.H{
		"post": post,
	})
}

func ClientIndex(c *gin.Context) {

	//Get the posts
	var posts []models.Client
	initializers.DB.Find(&posts)

	// Response with them
	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func ClientShow(c *gin.Context) {
	//Get id off url
	id := c.Param("id")

	//Get the posts
	var post models.Client
	initializers.DB.First(&post, id)

	// Response with them
	c.JSON(200, gin.H{
		"post": post,
	})
}

func ClientUpdate(c *gin.Context) {
	//Get the id off the url
	Id := c.Param("Id")

	//Get the data off req body
	var body struct {
		Email     string `gorm:"unique"`
		FirstName string
		LastName  string
		State     string
	}

	c.Bind(&body)

	//Find the post were upadating
	var post models.Client
	initializers.DB.First(&post, Id)

	//Updating it
	initializers.DB.Model(&post).Updates(models.Client{
		Email:     body.Email,
		FirstName: body.FirstName,
		LastName:  body.LastName,
		State:     body.State,
	})

	//Response with it
	c.JSON(200, gin.H{
		"post": post,
	})
}

func ClientDelete(c *gin.Context) {
	//Get the id off the url
	Id := c.Param("Id")

	//Delete the posts
	initializers.DB.Delete(&models.Client{}, Id)

	//Response
	c.Status(200)
}
