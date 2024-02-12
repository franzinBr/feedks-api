package handlers

import (
	"net/http"

	"github.com/franzinBr/feedks-api/api/dtos"
	"github.com/franzinBr/feedks-api/api/errors"
	"github.com/franzinBr/feedks-api/api/services"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service *services.UserService
}

func NewUserHandler() *UserHandler {
	return &UserHandler{
		service: services.NewUserService(),
	}
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	req := new(dtos.CreateUserRequest)

	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{
				"sucess":  false,
				"message": err.Error(),
			},
		)
		return
	}

	if err := h.service.CreateUser(req); err != nil {
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
		"message": "User created with sucess",
	})
}

func (h *UserHandler) Login(c *gin.Context) {
	req := new(dtos.LoginRequest)

	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{
				"sucess":  false,
				"message": err.Error(),
			},
		)
		return
	}

	tokenResponse, err := h.service.Login(req)

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
		"data":   tokenResponse,
	})

}
