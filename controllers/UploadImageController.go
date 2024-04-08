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
		BoardId     string
		Title       string
		StartTime   time.Time
		EndTime     time.Time
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

	plans := models.Plans{BoardId: body.BoardId, Title: body.Title, StartDate: body.StartTime, EndDate: body.EndTime, Description: body.Description}

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

// edit plan
func HandleEditPlan(c *gin.Context) {
	//Get the id of plan
	id := c.Param("id")

	// Get data off req body
	type ImageTab struct {
		Url string `json:"url"`
	}

	var body struct {
		Title       string
		StartTime   time.Time
		EndTime     time.Time
		Description string
		ImageUrls   []string
	}

	if c.BindJSON(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	// find the post were updating
	var plans models.Plans
	initializers.DB.Where("id =?", id).First(&plans)

	// updating it
	initializers.DB.Model(&plans).Updates(models.Plans{

		Title:       body.Title,
		StartDate:   body.StartTime,
		EndDate:     body.EndTime,
		Description: body.Description,
		ImageUrls:   body.ImageUrls,
	})

	c.JSON(200, gin.H{
		"plans": plans,
	})
}

// fuction gets all plans
func GetAllPlans(c *gin.Context) {
	var plans []models.Plans
	initializers.DB.Find(&plans)

	c.JSON(200, gin.H{
		"plans": plans,
	})
}

// fuction get plans for a specific BoardId
func GetPlansByBoardId(c *gin.Context) {
	boardId := c.Param("boardId")
	var plans []models.Plans
	initializers.DB.Where("board_id =?", boardId).Find(&plans)

	c.JSON(200, gin.H{
		"plans": plans,
	})
}
