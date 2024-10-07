package router

import (
	"api-service-go/config"
	"api-service-go/handler"
	"api-service-go/repository"
	"api-service-go/service"

	"github.com/gin-gonic/gin"
)


func AuthRouter(api *gin.RouterGroup) {
	authRepository := repository.NewAuthRepository(config.DB)
	authService := service.NewAuthRepository(authRepository)
	authHandler := handler.NewAuthHandler(authService)

	
	api.POST("/auth/register", authHandler.Register)
}