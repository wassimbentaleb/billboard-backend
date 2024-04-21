package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"

	"billboard/database"
	"billboard/entities"
)

type Plan struct {
	pg *database.Postgres
}

func NewPlan(pg *database.Postgres) *Plan {
	return &Plan{pg: pg}
}

func (plan *Plan) Create(c *gin.Context) {
	newPlan := entities.Plan{}
	err := c.BindJSON(&newPlan)
	if err != nil {
		log.Print(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := plan.pg.DB.Create(&newPlan)
	if result.Error != nil {
		log.Print(result.Error.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"plan": newPlan})
}

func (plan *Plan) Update(c *gin.Context) {
	// get the id from the url
	id := c.Param("id")

	// check if the plan exists
	var dbPlan entities.Plan
	plan.pg.DB.First(&dbPlan, id)
	if dbPlan.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "plan not found"})
		return
	}

	// retrieve data from req body and validate it
	updatedPlan := entities.Plan{}
	err := c.BindJSON(&updatedPlan)
	if err != nil {
		log.Print(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// update the plan in the database
	result := plan.pg.DB.Model(&updatedPlan).Clauses(clause.Returning{}).Where("id = ?", id).Updates(&updatedPlan)
	if result.Error != nil {
		log.Print(result.Error.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	// return updated plan
	c.JSON(http.StatusOK, gin.H{"plan": updatedPlan})
}

func (plan *Plan) FindAll(c *gin.Context) {
	var plans []entities.Plan

	// get all the plans
	plan.pg.DB.Find(&plans)

	// return all the plans
	c.JSON(200, gin.H{"plans": plans})
}

func (plan *Plan) FindByID(c *gin.Context) {
	// get the id from the url
	id := c.Param("id")

	// get the plan from the database by id
	var dbPlan entities.Plan
	plan.pg.DB.First(&dbPlan, id)
	if dbPlan.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "plan not found"})
		return
	}

	// return the plan
	c.JSON(http.StatusOK, gin.H{"plan": dbPlan})
}

func (plan *Plan) Delete(c *gin.Context) {
	id := c.Param("id")

	// check if the plan exists
	var dbPlan entities.Plan
	plan.pg.DB.First(&dbPlan, id)
	if dbPlan.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "plan not found"})
		return
	}

	// delete plan
	result := plan.pg.DB.Delete(&entities.Plan{}, id)
	if result.Error != nil {
		log.Print(result.Error.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.Status(http.StatusOK)
}
