package controllers

import (
	"billboard/database"
	"billboard/entities"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
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

	// prevent user creation
	newSubscription.User = nil

	result := subscription.pg.DB.Create(&newSubscription)
	if result.Error != nil {
		log.Print(result.Error.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	var client entities.User
	subscription.pg.DB.First(&client, newSubscription.UserID).Select("company_name")
	if client.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "client not found"})
		return
	}

	response := entities.SubscriptionResponse{
		ID:          newSubscription.ID,
		EndDate:     newSubscription.EndDate,
		CreatedDate: newSubscription.CreatedDate,
		Paid:        newSubscription.Paid,
		CompanyName: client.CompanyName,
	}

	c.JSON(http.StatusCreated, gin.H{"subscription": response})
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

func (subscription *Subscription) FindAll(c *gin.Context) {
	var subscriptions []entities.Subscription
	subscription.pg.DB.Preload("User").Find(&subscriptions)

	var response []entities.SubscriptionResponse
	for _, item := range subscriptions {
		response = append(response, entities.SubscriptionResponse{
			ID:          item.ID,
			EndDate:     item.EndDate,
			CreatedDate: item.CreatedDate,
			Paid:        item.Paid,
			CompanyName: item.User.CompanyName,
		})
	}

	c.JSON(http.StatusOK, gin.H{"subscriptions": response})
}

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

func (subscription *Subscription) FindByCompanyName(c *gin.Context) {
	value := c.Param("companyName")

	// check if the subscription exists
	var dbSubscription entities.Subscription
	subscription.pg.DB.Where("company_name = ?", value).First(&dbSubscription)
	if dbSubscription.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "subscription not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"subscription": dbSubscription})
}
