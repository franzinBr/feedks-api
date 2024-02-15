package api

import (
	"fmt"
	"log"
	"os"

	"github.com/franzinBr/feedks-api/api/middlewares"
	"github.com/franzinBr/feedks-api/api/routers"
	"github.com/gin-gonic/gin"
)

func InitServer() {
	r := gin.Default()

	r.Use(middlewares.Cors())

	registerRoutes(r)

	err := r.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))
	if err != nil {
		log.Fatal("Error on start API")
	}
}

func registerRoutes(r *gin.Engine) {
	api := r.Group("/api")

	v1 := api.Group("/v1")
	{
		health := v1.Group("/health")
		user := v1.Group("/user")
		feedback := v1.Group("/feedback", middlewares.Authentication())

		routers.Health(health)
		routers.User(user)
		routers.FeedBack(feedback)
	}
}
