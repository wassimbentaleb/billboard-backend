package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"

	"billboard/database"
	"billboard/entities"
)

type Board struct {
	pg *database.Postgres
}

func NewBoard(pg *database.Postgres) *Board {
	return &Board{pg: pg}
}

func (board *Board) Create(c *gin.Context) {
	// retrieve data from req body
	newBoard := entities.Board{}
	err := c.BindJSON(&newBoard)
	if err != nil {
		log.Print(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// create a new board
	result := board.pg.DB.Create(&newBoard)
	if result.Error != nil {
		log.Print(result.Error.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	// return it
	c.JSON(http.StatusCreated, gin.H{"board": newBoard})
}

func (board *Board) Update(c *gin.Context) {
	// get the id from the url
	id := c.Param("id")

	// check if the board exists
	var dbBoard entities.Board
	board.pg.DB.First(&dbBoard, id)
	if dbBoard.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "board not found"})
		return
	}

	// retrieve data from req body and validate it
	updatedBoard := entities.Board{}
	err := c.BindJSON(&updatedBoard)
	if err != nil {
		log.Print(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// update the board in the database
	result := board.pg.DB.Model(&updatedBoard).Clauses(clause.Returning{}).Where("id = ?", id).Updates(&updatedBoard)
	if result.Error != nil {
		log.Print(result.Error.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	// return updated board
	c.JSON(http.StatusOK, gin.H{"board": updatedBoard})
}

func (board *Board) FindAll(c *gin.Context) {
	var boards []entities.Board

	// get all the boards
	board.pg.DB.Find(&boards)

	// return all the boards
	c.JSON(200, gin.H{"boards": boards})
}

func (board *Board) FindByID(c *gin.Context) {
	// get the id from the url
	id := c.Param("id")

	// get the board from the database by id
	var dbBoard entities.Board
	board.pg.DB.First(&dbBoard, id)
	if dbBoard.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "board not found"})
		return
	}

	// return the board
	c.JSON(http.StatusOK, gin.H{"board": dbBoard})
}

func (board *Board) Delete(c *gin.Context) {
	id := c.Param("id")

	// check if the board exists
	var dbBoard entities.Board
	board.pg.DB.First(&dbBoard, id)
	if dbBoard.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "board not found"})
		return
	}

	// delete board
	result := board.pg.DB.Delete(&entities.Board{}, id)
	if result.Error != nil {
		log.Print(result.Error.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.Status(http.StatusOK)
}
