package controllers

import (
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
