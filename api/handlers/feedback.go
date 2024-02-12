package handlers

import (
	"net/http"

	"github.com/franzinBr/feedks-api/api/dtos"
	"github.com/franzinBr/feedks-api/api/errors"
	"github.com/franzinBr/feedks-api/api/services"
	"github.com/gin-gonic/gin"
)

type FeedBackHandler struct {
	service *services.FeedBackService
}

func NewFeedBackHandler() *FeedBackHandler {
	return &FeedBackHandler{
		service: services.NewFeedBackService(),
	}
}

func (h *FeedBackHandler) CreateFeedBack(c *gin.Context) {
	req := new(dtos.CreateFeedBackRequest)

	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{
				"sucess":  false,
				"message": err.Error(),
			},
		)
		return
	}

	userId := c.GetString("x-user-id")

	if err := h.service.CreateFeedBack(req, userId); err != nil {
		c.AbortWithStatusJSON(errors.GetStatusCodeFromError(err),
			gin.H{
				"sucess":  false,
				"message": err.Error(),
			},
		)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"sucess":  true,
		"message": "Feedback created with sucess",
	})
}

func (h *FeedBackHandler) ListFeedBacks(c *gin.Context) {
	req := new(dtos.PaginationRequest)

	if err := c.ShouldBind(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{
				"sucess":  false,
				"message": err.Error(),
			},
		)
		return
	}

	userId := c.GetString("x-user-id")

	feedbacks, err := h.service.ListFeedBacks(req, userId)

	if err != nil {
		c.AbortWithStatusJSON(errors.GetStatusCodeFromError(err),
			gin.H{
				"sucess":  false,
				"message": err.Error(),
			},
		)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"sucess": true,
		"data":   feedbacks,
	})

}