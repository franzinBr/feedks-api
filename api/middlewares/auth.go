package middlewares

import (
	"net/http"
	"os"
	"strings"

	"github.com/franzinBr/feedks-api/api/services"
	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	tokenService := services.NewTokenService()

	return func(c *gin.Context) {
		auth := strings.Split(c.GetHeader("Authorization"), " ")

		if len(auth) != 2 {
			abortUnauthorized(c)
			return
		}

		token := auth[1]

		claims, err := tokenService.GetClaims(token, os.Getenv("JWT_ACCESS_TOKEN_SECRET"))

		if err != nil {
			abortUnauthorized(c)
			return
		}

		c.Set("x-user-id", claims["id"])
		c.Next()
	}
}

func abortUnauthorized(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusUnauthorized,
		gin.H{
			"sucess":  false,
			"message": "Unauthorized",
		},
	)
}
