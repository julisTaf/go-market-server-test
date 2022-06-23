package controllers

import (
	"Go-market-test/pkg/database"
	u "Go-market-test/pkg/user"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Profile returns user data
func Profile(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	var user u.User

	id := c.Param("id")

	result := database.GlobalDB.Where("id = ?", id).First(&user)

	if result.Error == gorm.ErrRecordNotFound {
		c.JSON(404, gin.H{
			"msg": "user not found",
		})
		c.Abort()
		return
	}

	if result.Error != nil {
		c.JSON(500, gin.H{
			"msg": "could not get user profile",
		})
		c.Abort()
		return
	}

	user.Password = ""

	c.JSON(200, user)

	return
}
