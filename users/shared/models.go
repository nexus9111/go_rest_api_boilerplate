package shared

import (
	"github.com/gin-gonic/gin"
	"github.com/nexus9111/go-rest-api-boilerplate/utils"
)

type User struct {
	Id       string
	Email    string
	Username string
	Password string
}

func (u *User) Validate() bool {
	return u.Email != "" && u.Username != "" && u.Password != ""
}

func (u *User) Get(c *gin.Context) {
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
