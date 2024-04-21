package controllers

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type Upload struct{}

func NewUpload() *Upload {
	return &Upload{}
}

func (upload *Upload) FileUpload(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		log.Print(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	files := form.File["files"]
	urls := make([]string, 0, len(files))

	for _, file := range files {
		// generate a unique filename
		filename := strconv.FormatInt(time.Now().Unix(), 10) + "_" + strings.ReplaceAll(file.Filename, " ", "_")

		// save the file at static folder
		err = c.SaveUploadedFile(file, "static/"+filename)
		if err != nil {
			log.Print(err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// append the url to the urls slice
		serverStr := os.Getenv("SERVER")
		portStr := os.Getenv("PORT")
		urls = append(urls, fmt.Sprintf("%s:%s/media/%s", serverStr, portStr, filename))
	}

	c.JSON(http.StatusOK, gin.H{"media": urls})
}
