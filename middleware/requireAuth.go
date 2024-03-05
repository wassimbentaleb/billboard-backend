package middleware

import (
	"awesomeProject1/initializers"
	"awesomeProject1/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"os"
	"strings"
	"time"
)

func RequireAuth(c *gin.Context) {
	// Get the token from the request header
	tokenString := c.GetHeader("Authorization")

	// Check if token is present
	if tokenString == "" {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	// Remove "Bearer " prefix from token string
	tokenString = strings.Replace(tokenString, "Bearer ", "", 1)

	// Decode/validate the token

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(os.Getenv("SECRET")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Check the exp
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		//Get id off url
		id := c.Param("id")

		// Find the user with token sub
		var user models.User
		initializers.DB.First(&user, id)

		if user.ID == 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		// Attach to req
		c.Set("user", user)

		// Continue
		c.Next()

		fmt.Println(claims["foo"], claims["nbf"])
	} else {
		fmt.Println(err)
	}

}
