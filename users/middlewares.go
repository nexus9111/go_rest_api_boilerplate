package users

import (
	"GO_API_BOILERPLATE/utils"
	"github.com/gin-gonic/gin"
)

func IsAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user User
		user.get(c)
		c.Set("user", user)
		c.Next()
	}
}

func (u *User) get(c *gin.Context) {
	authorization := c.Request.Header["Authorization"]
	if len(authorization) == 0 {
		utils.ErrorResponse(c, 401, utils.ErrorUnauthorized)
		c.Abort()
	}

	// TODO: check token
	// TODO: get user from database

	u.Id = "1"
	u.Email = "test@test.com"
	u.Username = "test"
	u.Password = "test"
}
