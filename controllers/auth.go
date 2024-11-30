package controllers

import (
	"golang-comment/database"
	"golang-comment/models"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))
var validate = validator.New()

func Register(c *gin.Context) {
	var input models.User

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"data": nil,
			// "error":   err.Error(),
			"message": "Data not valid",
		})
		return
	}

	if err := validate.Struct(input); err != nil {
		errors := make(map[string]string)
		for _, err := range err.(validator.ValidationErrors) {
			errors[err.Field()] = err.Tag()
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"data":            nil,
			"validationError": errors,
			"message":         "Data not valid!",
		})
		return
	}

	var user models.User
	if err := database.DB.Where("email=?", input.Email).First(&user).Error; user.Email != "" && err == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"data": nil,
			// "error": err.Error(),
			"message": "Email already registered!",
		})
		return
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"data": nil,
			// "error":   err.Error(),
			"message": "Registered user failed!",
		})
		return
	}
	input.Password = string(hashPassword)
	database.DB.Create(&input)
	c.JSON(http.StatusOK, gin.H{
		"data":    nil,
		"message": "User registered successfully",
	})
}

func Login(c *gin.Context) {
	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"data":  nil,
			"error": err.Error(),
		})
		return
	}
	if err := validate.Var(input.Email, "required,email"); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"data":            nil,
			"validationError": err,
		})
		return
	}
	var user models.User
	if err := database.DB.Where("email=?", input.Email).First(&user).Error; user.Email == "" && err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"data": nil,
			// "error": err.Error(),
			"message": "wrong email or password!",
		})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"data": nil,
			// "error":   err.Error(),
			"message": "wrong email or password!!",
		})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": user.ID,
		"exp":    time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, _ := token.SignedString(jwtSecret)
	c.JSON(http.StatusOK, gin.H{
		"data":    tokenString,
		"message": "Login successfully",
	})
}
