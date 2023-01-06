package routers

import (
	"github.com/gin-gonic/gin"
	"main.go/handler"
)

func ClinicDomainRouter(app *gin.Engine) *gin.RouterGroup {
	clinicGroup := app.Group("/clinic")
	{
		clinicGroup.POST("", handler.SearchClinicbyLocation)
		clinicGroup.GET("", handler.GetAllDataClinic)
		clinicGroup.POST("/search", handler.SearchClinicUsingQuery)
	}
	return clinicGroup
}
