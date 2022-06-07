package routers

import (
	"github.com/gin-gonic/gin"
	"main.go/handler"
	"main.go/middleware"
)

func UserDomainRouter(app *gin.Engine) *gin.RouterGroup {
	userGroup := app.Group("/user") 
	{
		userGroup.POST("/register", handler.UserRegister)
		userGroup.POST("/login", handler.UserLogin)
		userGroup.GET("", middleware.AuthMiddleware(), handler.GetUser)
		userGroup.GET("/:id", handler.GetUserByID)
		userGroup.PATCH("/:id", handler.UpdateUser)
		userGroup.GET("/search", handler.SearchUserbyParams)
		userGroup.DELETE("/:id", handler.DeleteUser)
	}
	return userGroup
}