package routers

import (
	"github.com/gin-gonic/gin"
	"main.go/handler"
)

func DoctorDomainRouter(app *gin.Engine) *gin.RouterGroup {
	doctorGroup := app.Group("/doctor")
	{
		doctorGroup.POST("/register", handler.DoctorRegister)
		doctorGroup.POST("/login", handler.DoctorLogin)
		doctorGroup.GET("", handler.GetDoctor)
		doctorGroup.GET("/:id", handler.GetDoctorByID)
		doctorGroup.PATCH("/:id", handler.UpdateDataDoctor)
		doctorGroup.GET("/search", handler.SearchDoctorDatabyParams)
		doctorGroup.DELETE("/:id", handler.DeleteDoctorAccountUsingID)
	}
	return doctorGroup
}
