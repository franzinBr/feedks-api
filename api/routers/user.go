package routers

import (
	"github.com/franzinBr/feedks-api/api/handlers"
	"github.com/gin-gonic/gin"
)

func User(r *gin.RouterGroup) {
	h := handlers.NewUserHandler()

	r.POST("/", h.CreateUser)
	r.POST("/login", h.Login)
}
