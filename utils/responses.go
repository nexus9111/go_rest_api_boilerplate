package utils

import "github.com/gin-gonic/gin"

func ErrorResponse(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{
		"message": message,
	})
}

func SuccessResponse(c *gin.Context, code int, data map[string]any) {
	c.JSON(code, gin.H{
		"data": data,
	})
}

const (
	ErrorBadBody        string = "Invalid request body %s"
	ErrorInvalidRequest string = "Invalid request"
	ErrorUnauthorized   string = "Unauthorized"
	ErrorInternal       string = "Internal server error"
)
