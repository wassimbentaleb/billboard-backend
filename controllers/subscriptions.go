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
	// Check if CompanyName exists in users table
	var existingUser entities.User
	subscription.pg.DB.First(&existingUser, "company_name = ?", newSubscription.CompanyName)
	if existingUser.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "company name does not exist"})
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

// get all subscription
func (subscription *Subscription) FindAll(c *gin.Context) {
	var subscriptions []entities.Subscription
	subscription.pg.DB.Find(&subscriptions)

	c.JSON(http.StatusOK, gin.H{"subscriptions": subscriptions})
}

// get subscription by id
func (subscription *Subscription) FindByID(c *gin.Context) {
	id := c.Param("id")

	// check if the subscription exists
	var dbSubscription entities.Subscription
	subscription.pg.DB.First(&dbSubscription, id)
	if dbSubscription.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "subscription not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"subscription": dbSubscription})
}

// update subscription
func (subscription *Subscription) Update(c *gin.Context) {
	id := c.Param("id")

	// check if the subscription exists
	var dbSubscription entities.Subscription
	subscription.pg.DB.First(&dbSubscription, id)
	if dbSubscription.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "subscription not found"})
		return
	}

	updatedSubscription := entities.Subscription{}
	err := c.BindJSON(&updatedSubscription)
	if err != nil {
		log.Print(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := subscription.pg.DB.Save(&updatedSubscription)
	if result.Error != nil {
		log.Print(result.Error.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
}

// get subscription by company_name
func (subscription *Subscription) FindByCompanyName(c *gin.Context) {
	companyName := c.Param("company_name")

	// check if the subscription exists
	var dbSubscription entities.Subscription
	subscription.pg.DB.First(&dbSubscription, "company_name =?", companyName)
	if dbSubscription.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "subscription not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"subscription": dbSubscription})
}
