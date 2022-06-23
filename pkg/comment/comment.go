package comment

import (
	"Go-market-test/pkg/database"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
)

type Comment struct {
	gorm.Model
	Message    string `json:"message"`
	DealId     uint   `json:"deal_id"`
	ResponseTo uint   `json:"response_to"`
}

func NewComment(r *gin.Context) {
	var comment Comment
	var err error
	err = r.ShouldBindJSON(&comment)
	if err != nil {
		log.Println(err)

		r.JSON(400, gin.H{
			"msg": "invalid json",
		})
		r.Abort()

		return
	}

	result := database.GlobalDB.Create(&comment)
	if result.Error != nil {
		log.Println(result.Error)

		r.JSON(500, gin.H{
			"msg": "error creating user",
		})
		r.Abort()

		return
	}
	r.JSON(200, comment)
}

func GetComment(r *gin.Context) {
	var comments []Comment
	id := r.Param("id")
	result := database.GlobalDB.Find(&comments, "deal_id = ?", id)
	if result.Error != nil {
		r.JSON(500, gin.H{
			"msg": "Could not query deals",
		})
		r.Abort()
		return
	}
	r.JSON(200, comments)
}

func GetRestComment(r *gin.Context) {
	var comments []Comment
	id := r.Param("id")
	result := database.GlobalDB.Find(&comments, "response_to = ?", id)
	if result.Error != nil {
		r.JSON(500, gin.H{
			"msg": "Could not query deals",
		})
		r.Abort()
		return
	}
	r.JSON(200, comments)
}

func DeleteComment(r *gin.Context) {
	var comment Comment
	id := r.Param("id")
	result := database.GlobalDB.Delete(&comment, "id = ?", id)
	if result.Error != nil {
		r.JSON(400, gin.H{
			"msg": "invalid json",
		})
		r.Abort()

		return
	}
	r.JSON(200, true)
}
