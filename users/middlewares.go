package users

import (
	"GO_API_BOILERPLATE/users/shared"
	"github.com/gin-gonic/gin"
)

func IsAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user shared.User
		user.Get(c)
		c.Set("user", user)
		c.Next()
	}
}
