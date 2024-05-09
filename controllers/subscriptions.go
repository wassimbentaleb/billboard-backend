package controllers

import (
	"billboard/database"
	"billboard/entities"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Subscription struct {
	pg *database.Postgres
}

func NewSubscription(pg *database.Postgres) *Subscription {
	return &Subscription{pg: pg}
}

func (subscription *Subscription) Create(c *gin.Context) {
	newSubscription := entities.Subscription{}
	err := c.BindJSON(&newSubscription)
	if err != nil {
		log.Print(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := subscription.pg.DB.Create(&newSubscription)
	if result.Error != nil {
		log.Print(result.Error.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"subscription": newSubscription})
}

func (subscription *Subscription) Delete(c *gin.Context) {
	id := c.Param("id")

	// check if the subscription exists
	var dbSubscription entities.Subscription
	subscription.pg.DB.First(&dbSubscription, id)
	if dbSubscription.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "subscription not found"})
		return
	}

	// delete subscription
	result := subscription.pg.DB.Delete(&dbSubscription)
	if result.Error != nil {
		log.Print(result.Error.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.Status(http.StatusOK)
}
