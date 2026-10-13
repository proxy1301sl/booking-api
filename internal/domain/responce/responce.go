package responce

import "github.com/gin-gonic/gin"

type Response struct {
	Success bool   `json:"success"`
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
	Error   any    `json:"error,omitempty"`
	Data    any    `json:"data,omitempty"`
}

func SuccessResponse(c *gin.Context, StatusCode int, message string, data any) {
	c.JSON(StatusCode, Response{
		Success: true,
		Code:    StatusCode,
		Message: message,
		Data:    data,
	})
}

func ErrorResponse(c *gin.Context, statusCode int, massage string, err any) {
	c.JSON(statusCode, Response{
		Success: false,
		Code:    statusCode,
		Message: massage,
		Error:   err,
	})
}
