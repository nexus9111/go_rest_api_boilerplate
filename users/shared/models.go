package shared

import (
	"github.com/gin-gonic/gin"
	"github.com/nexus9111/go-rest-api-boilerplate/utils"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string `gorm:"index;unique"`
	Username string `gorm:"index;unique"`
	Password string
}

func (u *User) Validate() bool {
	return u.Email != "" && u.Username != "" && u.Password != ""
}

func (u *User) FindFromRequestToken(c *gin.Context) {
	authorization := c.Request.Header["Authorization"]
	if len(authorization) == 0 {
		utils.ErrorResponse(c, 401, utils.ErrorUnauthorized)
		c.Abort()
	}

	// TODO: check token
	// TODO: get user from database

	u.Email = "test@test.com"
	u.Username = "test"
	u.Password = "test"
}
