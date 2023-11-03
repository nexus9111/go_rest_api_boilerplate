package users

import (
	"GO_API_BOILERPLATE/users/shared"
	"GO_API_BOILERPLATE/utils"
	"fmt"
	"github.com/gin-gonic/gin"
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

	// TODO: Create user in database

	utils.SuccessResponse(c, 201, map[string]any{
		"email":    requestUser.Email,
		"username": requestUser.Username,
	})
}

func login(c *gin.Context) {
	type credentials struct {
		email    string
		password string
	}

	var requestCredentials credentials
	err := c.BindJSON(&requestCredentials)
	if err != nil || requestCredentials.email == "" || requestCredentials.password == "" {
		utils.ErrorResponse(c, 400, fmt.Sprintf(utils.ErrorBadBody, "[email, password]"))
		return
	}

	// TODO: Check credentials in database
	// TODO: Generate token
	token := "token"

	utils.SuccessResponse(c, 201, map[string]any{
		"token": token,
	})
}

func profile(c *gin.Context) {
	user, err := getUserFromContext(c)
	if err != nil {
		utils.ErrorResponse(c, 500, utils.ErrorInternal)
		return
	}

	utils.SuccessResponse(c, 200, map[string]any{
		"id":       user.Id,
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
		"id":       user.Id,
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
