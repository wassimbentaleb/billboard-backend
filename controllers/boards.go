package controllers

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"
	"net/http"
	"strings"

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

	// generate ref string on six hex characters
	// Create a byte slice to hold the random bytes
	bytes := make([]byte, 3) // 3 bytes = 6 hex characters

	// Read random bytes
	_, err = rand.Read(bytes)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Encode the bytes to a hexadecimal string
	hexString := strings.ToUpper(hex.EncodeToString(bytes))

	// Print the result
	fmt.Println("Random Hex String:", hexString)

	// Check if the hex string already exists in the database
	var dbBoard entities.Board
	board.pg.DB.First(&dbBoard, "ref =?", hexString)
	if dbBoard.ID != 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ref string already exists"})
		return
	}

	// set the ref string
	newBoard.Ref = hexString

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
	result := board.pg.DB.Select("Plans").Delete(&entities.Board{ID: dbBoard.ID})
	if result.Error != nil {
		log.Print(result.Error.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.Status(http.StatusOK)
}

func (board *Board) linkBoardUpdate(c *gin.Context) {

	//check if the ref exists dans la board table
	var dbBoard entities.Board
	board.pg.DB.First(&dbBoard, "ref =?", c.Param("ref"))
	if dbBoard.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "board not found"})
		return
	}

	// retrieve data from req body and validate it
	updatedBoard := entities.Board{
		Name:        dbBoard.Name,
		Description: dbBoard.Description,
		Address:     dbBoard.Address,
	}
	err := c.BindJSON(&updatedBoard)
	if err != nil {
		log.Print(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// update the board in the database
	result := board.pg.DB.Model(&updatedBoard).Clauses(clause.Returning{}).Where("id =?", dbBoard.ID).Updates(&updatedBoard)
	if result.Error != nil {
		log.Print(result.Error.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	// return updated board
	c.JSON(http.StatusOK, gin.H{"board": updatedBoard})

}
