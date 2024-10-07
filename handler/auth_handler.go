package handler

import (
	"api-service-go/dto"
	errorhandler "api-service-go/errorHandler"
	"api-service-go/helper"
	"api-service-go/service"
	"net/http"

	"github.com/gin-gonic/gin"
)


type authHandler struct {
service service.AuthService
}

func NewAuthHandler(s service.AuthService) *authHandler{
	return &authHandler{
		service: s,
	}
}


func (h *authHandler) Register(c *gin.Context) {
	var register dto.RegisterRequest


	if err := c.ShouldBindJSON(&register); err != nil {
		errorhandler.HandlerError(c,&errorhandler.BadRequestError{
			Message: err.Error(),
		})
		return
	}

	if err := h.service.Register(&register); err != nil {
		errorhandler.HandlerError(c, err)
		return
	}

	res := helper.Response(dto.ResponseParams{
		StatusCode: http.StatusCreated,
		Message: "Register successfully",
	})
	c.JSON(http.StatusCreated,res)
}