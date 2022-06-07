package routers

import (
	"github.com/gin-gonic/gin"
	"main.go/cors"
)

func InitGin() *gin.Engine {
	app := gin.Default()
	app.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})
	app.Use(cors.CORSPreflightMiddleware())
	ArticleDomainRouter(app)
	ClinicDomainRouter(app)
	DoctorDomainRouter(app)
	HospitalDomainRouter(app)
	UserDomainRouter(app)
	return app
}
