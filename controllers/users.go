package controllers

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"

	"billboard/database"
	"billboard/entities"
)

type User struct {
	pg *database.Postgres
}

func NewUser(pg *database.Postgres) *User {
	return &User{pg: pg}
}

func (user *User) Signup(c *gin.Context) {
	newUser := entities.User{}
	err := c.BindJSON(&newUser)
	if err != nil {
		log.Print(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// check if email is already in use
	var existingUser entities.User
	user.pg.DB.First(&existingUser, "email = ?", newUser.Email)
	if existingUser.ID != 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "email already in use"})
		return
	}

	// hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), 10)
	if err != nil {
		log.Print(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	newUser.Password = string(hash)

	// create the user
	result := user.pg.DB.Create(&newUser)
	if result.Error != nil {
		log.Print(result.Error.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	// generate a jwt token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": newUser.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	// sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		log.Print(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// send it back
	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

func (user *User) Login(c *gin.Context) {
	loginReq := entities.LoginRequest{}
	err := c.BindJSON(&loginReq)
	if err != nil {
		log.Print(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// look up requested user
	var dbUser entities.User
	user.pg.DB.First(&dbUser, "email = ?", loginReq.Email)
	if dbUser.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid email address"})
		return
	}

	// compare request password with real user password using hashs
	err = bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(loginReq.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid password"})
		return
	}

	// generate a jwt token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": dbUser.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	// sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		log.Print(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// send it back
	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}
