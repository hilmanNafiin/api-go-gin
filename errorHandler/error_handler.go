package errorhandler

import (
	"api-service-go/dto"
	"api-service-go/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandlerError(c *gin.Context, err error) {
	var statusCode int


	switch err.(type) {
	case *NotFoundError:
		statusCode = http.StatusNotFound
	case *BadRequestError:
		statusCode = http.StatusBadRequest
	case *UnauthorizedError:
		statusCode = http.StatusUnauthorized
	case *InternalServerError:
		statusCode = http.StatusInternalServerError
	}

	response := helper.Response(dto.ResponseParams{
		StatusCode: statusCode,
		Message: err.Error(),
	})

	c.JSON(statusCode, response)
	
}