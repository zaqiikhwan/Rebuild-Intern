package routers

import (
	"github.com/gin-gonic/gin"
	"main.go/cors"
)

func InitGin() *gin.Engine {
	app := gin.Default()
	app.Use(cors.CORSPreflightMiddleware())
	ArticleDomainRouter(app)
	ClinicDomainRouter(app)
	DoctorDomainRouter(app)
	HospitalDomainRouter(app)
	UserDomainRouter(app)
	return app
}