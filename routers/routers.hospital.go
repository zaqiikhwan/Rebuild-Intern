package routers

import (
	"github.com/gin-gonic/gin"
	"main.go/handler"
)

func HospitalDomainRouter(app *gin.Engine) *gin.RouterGroup {
	hospitalGroup := app.Group("/hospital")
	{
		hospitalGroup.POST("/register", handler.HospitalRegister)
		hospitalGroup.GET("", handler.GetAllHospitalList)
		hospitalGroup.GET("/:id", handler.SearchHospitalByIDQuery)
		hospitalGroup.PATCH("/:id", handler.UpdateHospitalDatabyID)
		hospitalGroup.GET("/search", handler.SearchHospitalbyCityQuery)
	}
	return hospitalGroup
}
