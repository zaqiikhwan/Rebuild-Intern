package routers

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func InitGin() *gin.Engine {
	// change to release mode
	gin.SetMode(gin.ReleaseMode)
	app := gin.Default()
	// configuration CORS
	app.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", os.Getenv("HOST"))
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		if c.Request.Method == "OPTIONS" {
			c.Writer.Header().Set("Content-Type", "application/json")
			c.AbortWithStatus(204)
		} else {
			c.Next()
		}
	})
	app.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "success",
			"text":    "api connected",
		})
	})
	ArticleDomainRouter(app)
	ClinicDomainRouter(app)
	DoctorDomainRouter(app)
	HospitalDomainRouter(app)
	UserDomainRouter(app)
	return app
}
