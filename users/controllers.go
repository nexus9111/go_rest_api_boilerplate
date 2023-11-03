package users

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/nexus9111/go-rest-api-boilerplate/config/database"
	"github.com/nexus9111/go-rest-api-boilerplate/users/shared"
	"github.com/nexus9111/go-rest-api-boilerplate/utils"
)

func getUserFromContext(c *gin.Context) (shared.User, error) {
	value, exists := c.Get("user")
	if !exists {
		return shared.User{}, fmt.Errorf("user not found in context")
	}

	user, ok := value.(shared.User)
	if !ok {
		return shared.User{}, fmt.Errorf("user is not of type User")
	}

	return user, nil
}

func create(c *gin.Context) {
	var requestUser shared.User
	err := c.BindJSON(&requestUser)
	if err != nil || !requestUser.Validate() {
		utils.ErrorResponse(c, 400, fmt.Sprintf(utils.ErrorBadBody, "[email, username, password]"))
		return
	}

	result := database.Context.DB.Create(&requestUser)
	if result.Error != nil {
		utils.ErrorResponse(c, 400, utils.AlreadyUsedEmailorPassword)
		return
	}

	utils.SuccessResponse(c, 201, map[string]any{
		"id":       requestUser.ID,
		"email":    requestUser.Email,
		"username": requestUser.Username,
	})
}

func login(c *gin.Context) {
	type Credentials struct {
		Email    string
		Password string
	}

	var requestCredentials Credentials
	err := c.BindJSON(&requestCredentials)

	if err != nil || requestCredentials.Email == "" || requestCredentials.Password == "" {
		utils.ErrorResponse(c, 400, fmt.Sprintf(utils.ErrorBadBody, "[email, password]"))
		return
	}

	var user shared.User
	findUser := database.Context.DB.First(&user, "email = ? AND password = ?", requestCredentials.Email, requestCredentials.Password)
	if findUser.Error != nil {
		utils.ErrorResponse(c, 401, utils.ErrorUnauthorized)
		return
	}

	// TODO: Generate token
	token := "token"

	utils.SuccessResponse(c, 201, map[string]any{
		"token":    token,
		"username": user.Username,
		"ID":       user.ID,
	})
}

func profile(c *gin.Context) {
	user, err := getUserFromContext(c)
	if err != nil {
		utils.ErrorResponse(c, 500, utils.ErrorInternal)
		return
	}

	utils.SuccessResponse(c, 200, map[string]any{
		"id":       user.ID,
		"email":    user.Email,
		"username": user.Username,
		"token":    c.Request.Header["Authorization"],
	})
}

func update(c *gin.Context) {
	user, err := getUserFromContext(c)
	if err != nil {
		utils.ErrorResponse(c, 500, utils.ErrorInternal)
		return
	}

	// TODO: Update user in database

	utils.SuccessResponse(c, 200, map[string]any{
		"id":       user.ID,
		"email":    user.Email,
		"username": user.Username,
		"token":    c.Request.Header["Authorization"],
	})
}

func deleteOne(c *gin.Context) {
	user, err := getUserFromContext(c)
	if err != nil {
		utils.ErrorResponse(c, 500, utils.ErrorInternal)
		return
	}

	// TODO: Delete user in database
	_ = user

	utils.SuccessResponse(c, 200, map[string]any{})
}
