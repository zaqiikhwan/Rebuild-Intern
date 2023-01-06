package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"main.go/database"
	"main.go/domain"
)

func SearchClinicbyLocation(c *gin.Context) {
	var body domain.Scrape
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Location is invalid.",
			"error":   err.Error(),
		})
		return
	}
	var queryResults []domain.Scrape
	trx := database.GetDB()
	if result := trx.Where("Location = ?", body.Location).Find(&queryResults); result.Error != nil {
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

func GetAllDataClinic(c *gin.Context) {
	var queryResults []domain.Scrape
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

func SearchClinicUsingQuery(c *gin.Context) {
	location, isLocationExist := c.GetQuery("location")

	var queryResults []domain.Scrape
	db := database.GetDB()

	if isLocationExist {
		search := "%" + location + "%"
		db = db.Where("location ILIKE ?", search)
	}

	if result := db.Find(&queryResults); result.Error != nil {
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
		"search": location,
	})
}