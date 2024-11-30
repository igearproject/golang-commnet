package controllers

import (
	"golang-comment/database"
	"golang-comment/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func GetComments(c *gin.Context) {
	var comments []models.DataGetComment
	database.DB.Preload("User").Table("comments").Find(&comments)
	c.JSON(http.StatusOK, gin.H{
		"data": comments,
	})
}

func CreateComment(c *gin.Context) {
	var input models.DataGetComment
	if userId, exist := c.Get("userId"); !exist {
		c.JSON(http.StatusInternalServerError, gin.H{
			"data":    nil,
			"message": "User id not found",
		})
		return
	} else {
		input.UserId = userId.(uint)
	}
	// fmt.Println(input.UserId)
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"data": nil,
			// "error": err.Error(),
			"message": "data not valid",
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
			"message":         "data not valid",
		})
		return
	}
	database.DB.Table("comments").Create(&input)
	c.JSON(http.StatusOK, gin.H{
		"data":    input,
		"message": "Create new comment successfully",
	})
}

func UpdateComment(c *gin.Context) {
	var comment models.Comment
	var input models.DataGetComment
	id := c.Param("id")

	if userId, exist := c.Get("userId"); !exist {
		c.JSON(http.StatusInternalServerError, gin.H{
			"data":    nil,
			"message": "User id not found",
		})
		return
	} else {
		input.UserId = userId.(uint)
	}

	if err := database.DB.First(&comment, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"data":    nil,
			"message": "Comment not found",
		})
		return
	}

	if comment.UserId != input.UserId {
		c.JSON(http.StatusForbidden, gin.H{
			"data":    nil,
			"message": "you don't have access to update this comment",
		})
		return
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"data": nil,
			// "error": err.Error(),
			"message": "data not valid",
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
			"message":         "data not valid",
		})
		return
	}

	comment.Title = input.Title
	comment.Body = input.Body

	database.DB.Save(&comment)

	c.JSON(http.StatusOK, gin.H{
		"data":    comment,
		"message": "Comment updated!",
	})
}

func DeleteComment(c *gin.Context) {
	var comment models.Comment
	id := c.Param("id")

	userId, exist := c.Get("userId")
	if !exist {
		c.JSON(http.StatusInternalServerError, gin.H{
			"data":    nil,
			"message": "User id not found",
		})
		return
	}
	if err := database.DB.Where("id=? AND user_id=?", id, userId).First(&comment).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"data":  nil,
			"error": "Comment not found or you don't have access",
		})
		return
	}

	database.DB.Delete(&comment)
	c.JSON(http.StatusOK, gin.H{
		"data":    nil,
		"message": "Comment deleted!",
	})
}
