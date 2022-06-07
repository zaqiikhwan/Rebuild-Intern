package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"main.go/database"
	"main.go/domain"
)

func HospitalRegister(c *gin.Context) {
	var body domain.Hospital
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Body is invalid.",
			"error":   err.Error(),
		})
		return
	}
	hospital := domain.Hospital{
		NameCity:     body.NameCity,
		HospitalName: body.HospitalName,
		Contact:      body.Contact,
		Address:      body.Address,
		Link:         body.Link,
	}
	if result := database.GetDB().Create(&hospital); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Error when inserting into the database.",
			"error":   result.Error.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"message": "Hospital registered successfully.",
		"data": gin.H{
			"id": hospital.ID,
		},
	})
}

func SearchHospitalByIDQuery(c *gin.Context) {
	id, isIdExists := c.Params.Get("id")
	if !isIdExists {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "ID is not supplied.",
		})
		return
	}
	hospital := domain.Hospital{}
	if result := database.GetDB().Where("id = ?", id).Take(&hospital); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Error when querying the database.",
			"error":   result.Error.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Query successful.",
		"data":    hospital,
	})
}

func GetAllHospitalList(c *gin.Context) {
	var queryResults []domain.Hospital
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
		"data": gin.H{
			"result": queryResults,
		},
	})
}

func SearchHospitalbyCityQuery(c *gin.Context) {
	namakota, isNamaKotaExists := c.GetQuery("name_city")
	if !isNamaKotaExists {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Query is not supplied.",
		})
		return
	}

	var queryResults []domain.Hospital
	trx := database.GetDB()
	if isNamaKotaExists {
		trx = trx.Where("name_city LIKE ?", "%"+namakota+"%")
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
		"data": gin.H{
			"result": queryResults,
		},
	})
}

func UpdateHospitalDatabyID(c *gin.Context) {
	id, isIdExists := c.Params.Get("id")
	if !isIdExists {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "ID is not supplied.",
		})
		return
	}
	var body domain.Hospital
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Body is invalid.",
			"error":   err.Error(),
		})
		return
	}
	parsedId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "ID is invalid.",
			"error":   err.Error(),
		})
		return
	}
	hospital := domain.Hospital{
		ID:           uint(parsedId),
		NameCity:     body.NameCity,
		HospitalName: body.HospitalName,
		Contact:      body.Contact,
		Address:      body.Address,
		Link:         body.Link,
	}
	result := database.GetDB().Model(&hospital).Updates(hospital)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Error when updating the database.",
			"error":   result.Error.Error(),
		})
		return
	}
	if result = database.GetDB().Where("id = ?", parsedId).Take(&hospital); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Error when querying the database.",
			"error":   result.Error.Error(),
		})
		return
	}
	if result.RowsAffected < 1 {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"message": "Hospital not found.",
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"message": "Update successful.",
		"data":    hospital,
	})
}
