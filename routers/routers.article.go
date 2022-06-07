package routers

import (
	"github.com/gin-gonic/gin"
	"main.go/handler"
)

func ArticleDomainRouter(app *gin.Engine) *gin.RouterGroup {
	articleGroup := app.Group("/article")
	{
		articleGroup.POST("", handler.GetArticleByID)
		articleGroup.GET("", handler.GetAllArticle)
		articleGroup.POST("/category", handler.GetArticleByCategory)
		articleGroup.GET("/search", handler.SearchArticleUsingQuery)
	}
	return articleGroup
}
