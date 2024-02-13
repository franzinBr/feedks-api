package routers

import (
	"github.com/franzinBr/feedks-api/api/handlers"
	"github.com/gin-gonic/gin"
)

func FeedBack(r *gin.RouterGroup) {
	h := handlers.NewFeedBackHandler()

	r.POST("/", h.CreateFeedBack)
	r.GET("/", h.ListFeedBacks)
	r.DELETE("/:id", h.DeleteFeedback)
}
