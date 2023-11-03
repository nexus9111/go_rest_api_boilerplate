package users

import (
	"github.com/gin-gonic/gin"
	"github.com/nexus9111/go-rest-api-boilerplate/users/shared"
)

func IsAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user shared.User
		user.Get(c)
		c.Set("user", user)
		c.Next()
	}
}
