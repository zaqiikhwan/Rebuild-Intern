package routers

import (
	"github.com/gin-gonic/gin"
	"main.go/cors"
)

func InitGin() *gin.Engine {
	app := gin.Default()
	app.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
<<<<<<< HEAD
			"message": "Hello World!",
=======
			"Text": "Hello world",
>>>>>>> 091b49065a66f2f122b6fa7b134b8a8e89f632bf
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
