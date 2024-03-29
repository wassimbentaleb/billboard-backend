package controllers

import (
	"awesomeProject1/initializers"
	"awesomeProject1/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

func HandleFileUpload(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("error prasing: %s", err.Error()))
		return
	}

	files := form.File["files"]
	var urls []string

	for _, file := range files {
		//save the file
		filename := strconv.FormatInt(time.Now().Unix(), 10) + file.Filename
		err = c.SaveUploadedFile(file, "static/"+filename)
		if err != nil {
			c.JSON(400, gin.H{"error": "Failed to upload image"})
			return
		}
		urls = append(urls, fmt.Sprintf("http://localhost:5000/media/%s", filename))
	}

	c.JSON(200, gin.H{"media": urls})
}

func HandleAddPlan(c *gin.Context) {
	// Get data off req body
	type ImageTab struct {
		Url string `json:"url"`
	}

	var body struct {
		Title       string
		StartDate   time.Time
		EndDate     time.Time
		Description string
		ImageUrls   []string
	}

	if c.BindJSON(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}
	// Create a Plan

	plans := models.Plans{Title: body.Title, StartDate: body.StartDate, EndDate: body.EndDate, Description: body.Description}

	for _, image := range body.ImageUrls {
		plans.ImageUrls = append(plans.ImageUrls, image)
	}

	result := initializers.DB.Create(&plans)

	if result.Error != nil {
		c.JSON(400, gin.H{"err": result.Error})
		return
	}

	c.JSON(200, gin.H{
		"plans": plans,
	})
}
