package controllers

import (
	"log"
	"net/http"
	"os"
	"time"

	"billboard/database"
	"billboard/entities"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	pg *database.Postgres
}

func NewUser(pg *database.Postgres) *User {
	return &User{pg: pg}
}

func (user *User) Create(c *gin.Context) {
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
	const DefaultPassword = "12345678"
	hash, err := bcrypt.GenerateFromPassword([]byte(DefaultPassword), 10)
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

	// return status created
	c.Status(http.StatusCreated)
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

func (user *User) FindAll(c *gin.Context) {
	var users []entities.User

	// get all the users
	user.pg.DB.Where("is_admin = ?", false).Find(&users)
	c.JSON(http.StatusOK, gin.H{"users": users})
}

func (user *User) FindCurrent(c *gin.Context) {
	currentUser, _ := c.Get("user")
	id := currentUser.(entities.User).ID

	// get the user from the database by id
	var dbUser entities.User
	user.pg.DB.First(&dbUser, id)
	if dbUser.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": dbUser})
}

func (user *User) Delete(c *gin.Context) {
	id := c.Param("id")

	// check if the user exists
	var dbUser entities.User
	user.pg.DB.First(&dbUser, id)
	if dbUser.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	// delete user
	result := user.pg.DB.Select("Subscriptions").Delete(&entities.User{ID: dbUser.ID})
	if result.Error != nil {
		log.Print(result.Error.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.Status(http.StatusOK)
}

func (user *User) Update(c *gin.Context) {
	//get the id from the url
	id := c.Param("id")

	// check if the user exists
	var dbUser entities.User
	user.pg.DB.First(&dbUser, id)
	if dbUser.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	// retrieve data from req body and validate it
	updatedUser := entities.UpdateUserRequest{}
	err := c.BindJSON(&updatedUser)
	if err != nil {
		log.Print(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// update password if provided
	if updatedUser.Password != "" {
		// compare request password with real user password using hashs
		err = bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(updatedUser.Password))
		if err != nil {
			log.Print("invalid password")
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid password"})
			return
		}

		// test if new password is provided
		if updatedUser.NewPassword == "" {
			log.Print("new password is required")
			c.JSON(http.StatusBadRequest, gin.H{"error": "new password is required"})
			return
		}

		// check the minimum new password length
		if len(updatedUser.NewPassword) < 8 {
			log.Print("new password must be at least 8 characters long")
			c.JSON(http.StatusBadRequest, gin.H{"error": "new password must be at least 8 characters long"})
			return
		}

		// hash the new password
		hash, err := bcrypt.GenerateFromPassword([]byte(updatedUser.NewPassword), 10)
		if err != nil {
			log.Print(err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		updatedUser.Password = string(hash)
	}

	// update user
	result := user.pg.DB.Model(&dbUser).Updates(updatedUser)
	if result.Error != nil {
		log.Print(result.Error.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.Status(http.StatusOK)
}
