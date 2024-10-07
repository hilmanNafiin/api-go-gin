package main

import (
	"api-service-go/config"
	"api-service-go/router"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {


	config.LoadConfig()
	config.LoadDB()

	
	
	r := gin.Default()


	api := r.Group("/api/v1")

	router.AuthRouter(api)

	r.Run(fmt.Sprintf(":%s", config.ENV.PORT))
}
