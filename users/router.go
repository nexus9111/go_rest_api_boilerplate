package users

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	fmt.Println("Initializing users router...")
	v1 := r.Group("/v1/users")
	{
		v1.POST("/", create)
		v1.POST("/login", login)
		v1.GET("/", IsAuth(), profile)
		v1.PUT("/", IsAuth(), update)
		v1.DELETE("/", IsAuth(), deleteOne)
	}
}
