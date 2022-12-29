package handler

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"main.go/database"
	"main.go/domain"
)

func DoctorRegister(c *gin.Context) {
	// var body postRegisterBody
	var body domain.User
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Body is invalid.",
			"error":   err.Error(),
		})
		return
	}
	doctor := domain.Doctor{
		Name:     body.Name,
		Email:    body.Email,
		Password: body.Password,
		Username: body.Username,
	}
	if result := database.GetDB().Create(&doctor); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Error when inserting into the database.",
			"error":   result.Error.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"message": "User Registered successfully",
		"data": gin.H{
			"id": doctor.ID,
		},
	})
}

func DoctorLogin(c *gin.Context) {
	// var body postLoginBody
	var body domain.Doctor
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Body is invalid.",
			"error":   err.Error(),
		})
		return
	}
	doctor := domain.Doctor{}
	if result := database.GetDB().Where("email = ? ", body.Email).Take(&doctor); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Error when querying the database.",
			"error":   result.Error.Error(),
		})
		return
	}
	if doctor.Password == body.Password {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"id":  doctor.ID,
			"exp": time.Now().Add(time.Hour * 24 * 7).Unix(),
		})
		tokenString, err := token.SignedString([]byte("passwordBuatSigning"))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": "Error when generating the token.",
				"error":   err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "Password is correct.",
			"data": gin.H{
				"id":       doctor.ID,
				"name":     doctor.Name,
				"username": doctor.Username,
				"token":    tokenString,
			},
		})
		return
	} else {
		c.JSON(http.StatusForbidden, gin.H{
			"success": false,
			"message": "Password is incorrect.",
		})
		return
	}
}

func GetDoctor(c *gin.Context) {
	// id, _ := c.Get("id")
	var doctors []domain.Doctor
	if result := database.GetDB().Find(&doctors); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Error when querying the database.",
			"error":   result.Error.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Query successful",
		"data":    doctors,
	})
}

func GetDoctorByID(c *gin.Context) {
	id, isIdExists := c.Params.Get("id")
	if !isIdExists {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "ID is not supplied.",
		})
		return
	}
	doctor := domain.Doctor{}
	if result := database.GetDB().Where("id = ?", id).Take(&doctor); result.Error != nil {
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
		"data":    doctor,
	})
}

func UpdateDataDoctor(c *gin.Context) {
	id, isIdExists := c.Params.Get("id")
	if !isIdExists {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "ID is not supplied.",
		})
		return
	}
	var body domain.Doctor
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
	doctor := domain.Doctor{
		ID:       uint(parsedId),
		Name:     body.Name,
		Email:    body.Email,
		Password: body.Password,
		Username: body.Username,
	}
	result := database.GetDB().Model(&doctor).Updates(doctor)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Error when updating the database.",
			"error":   result.Error.Error(),
		})
		return
	}
	if result = database.GetDB().Where("id = ?", parsedId).Take(&doctor); result.Error != nil {
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
			"message": "User not found.",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Update successful.",
		"data":    doctor,
	})
}

func SearchDoctorDatabyParams(c *gin.Context) {
	name, isNameExists := c.GetQuery("name")
	email, isEmailExists := c.GetQuery("email")
	username, isUsernameExists := c.GetQuery("username")
	if !isNameExists && !isEmailExists && !isUsernameExists {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Query is not supplied.",
		})
		return
	}

	var queryResults []domain.Doctor
	trx := database.GetDB()
	if isNameExists {
		trx = trx.Where("name LIKE ?", "%"+name+"%")
	}
	if isEmailExists {
		trx = trx.Where("email LIKE ?", "%"+email+"%")
	}
	if isUsernameExists {
		trx = trx.Where("username LIKE ?", "%"+username+"%")
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
			"query": gin.H{
				"name":     name,
				"email":    email,
				"username": username,
			},
			"result": queryResults,
		},
	})
}

func DeleteDoctorAccountUsingID(c *gin.Context) {
	id, isIdExists := c.Params.Get("id")
	if !isIdExists {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "ID is not supplied.",
		})
		return
	}
	parsedId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "ID is invalid.",
		})
		return
	}
	doctor := domain.Doctor{
		ID: uint(parsedId),
	}
	if result := database.GetDB().Delete(&doctor); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Error when deleting from the database.",
			"error":   result.Error.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Delete successful.",
	})
}
