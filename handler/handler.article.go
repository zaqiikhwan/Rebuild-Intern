package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"main.go/database"
	"main.go/domain"
)

func GetAllArticle(c *gin.Context) {
	var queryResults []domain.Article
	trx := database.GetDB()
	if result := trx.Find(&queryResults); result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Query is not supplied.",
			"error":   result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Search successful",
		"data":    queryResults,
	})
}

func GetArticleByID(c *gin.Context) {
	var body domain.Article
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Category is invalid.",
			"error":   err.Error(),
		})
		return
	}
	var queryResults []domain.Article
	trx := database.GetDB()
	if result := trx.Where("ID = ?", body.ID).Find(&queryResults); result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Query is not supplied.",
			"error":   result.Error.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Search successful",
		"data":    queryResults,
	})
}

func GetArticleByCategory(c *gin.Context) {
	var body domain.Article
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Category is invalid.",
			"error":   err.Error(),
		})
		return
	}
	var queryResults []domain.Article
	trx := database.GetDB()
	if result := trx.Where("Category = ?", body.Category).Find(&queryResults); result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Query is not supplied.",
			"error":   result.Error.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Search successful",
		"data":    queryResults,
	})
}

func SearchArticleUsingQuery(c *gin.Context) {
	trx := database.GetDB()
	
	title, isTitleExists := c.GetQuery("title")
	category, isCategoryExists := c.GetQuery("category")

	var queryResults []domain.Article
	
	if isTitleExists {
		trx = trx.Where("title LIKE ?", "%"+title+"%")
	}
	if isCategoryExists {
		trx = trx.Where("category LIKE ?", "%"+category+"%")
	}

	if result := trx.Find(&queryResults); result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Query is not supplied.",
			"error":   result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Search successful",
		"data":    queryResults,
	})
}