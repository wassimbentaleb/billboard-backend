package middleware

import (
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"

	"billboard/database"
	"billboard/entities"
)

type AuthMiddleware struct {
	pg *database.Postgres
}

func NewAuthMiddleware(pg *database.Postgres) *AuthMiddleware {
	return &AuthMiddleware{pg: pg}
}

func (middleware *AuthMiddleware) RequireAuth(c *gin.Context) {
	// skip auth for login and signup
	if strings.HasPrefix(c.Request.URL.Path, "/auth") {
		c.Next()
		return
	}

	// get the token from the request header
	tokenString := c.GetHeader("Authorization")

	// check if token is present
	if tokenString == "" {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	// remove "Bearer " prefix from token string
	tokenString = strings.Replace(tokenString, "Bearer ", "", 1)

	// decode and validate the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		log.Print(err.Error())
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// check the expiration time
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		// check user exists
		var user entities.User
		id := claims["sub"]
		middleware.pg.DB.First(&user, id)
		if user.ID == 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		// attach to req
		c.Set("user", user)

		// continue
		c.Next()
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
}
