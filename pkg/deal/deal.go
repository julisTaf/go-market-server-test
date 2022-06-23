package deal

import (
	"Go-market-test/pkg/database"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
)

type Deal struct {
	gorm.Model
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Author      string  `json:"author"`
	AuthorName  string  `json:"author_name"`
	Image       string  `json:"image"`
}

func NewDeal(r *gin.Context) {
	var deal Deal
	var err error
	err = r.ShouldBindJSON(&deal)
	if err != nil {
		log.Println(err)

		r.JSON(400, gin.H{
			"msg": "invalid json",
		})
		r.Abort()

		return
	}

	result := database.GlobalDB.Create(&deal)
	if result.Error != nil {
		log.Println(result.Error)

		r.JSON(500, gin.H{
			"msg": "error creating user",
		})
		r.Abort()

		return
	}
	r.JSON(200, deal)
}

func GetDeal(r *gin.Context) {
	var deal Deal
	id := r.Param("id")
	result := database.GlobalDB.Where("id=?", id).First(&deal)
	if result.Error == gorm.ErrRecordNotFound {
		r.JSON(404, gin.H{
			"msg": "user not found",
		})
		r.Abort()
		return
	}

	if result.Error != nil {
		r.JSON(500, gin.H{
			"msg": "could not get deal",
		})
		r.Abort()
		return
	}

	r.JSON(200, deal)

	return
}

func GetAllDeals(r *gin.Context) {
	var deals []Deal
	result := database.GlobalDB.Find(&deals)
	if result.Error != nil {
		r.JSON(500, gin.H{
			"msg": "Could not query deals",
		})
		r.Abort()
		return
	}
	r.JSON(200, deals)
}

func GetUserDeals(r *gin.Context) {
	var deals []Deal
	author := r.Param("id")
	result := database.GlobalDB.Find(&deals, "author = ?", author)
	if result.Error != nil {
		r.JSON(500, gin.H{
			"msg": "Something went wrong",
		})
		r.Abort()
		return
	}
	r.JSON(200, deals)
}

func UpdateDeal(r *gin.Context) {
	var dealA Deal
	var err error
	err = r.ShouldBindJSON(&dealA)
	if err != nil {
		log.Println(err)

		r.JSON(400, gin.H{
			"msg": "invalid json",
		})
		r.Abort()

		return
	}

	var deal Deal
	result := database.GlobalDB.Model(&deal).Where("id=?", dealA.ID).Updates(map[string]interface{}{"name": dealA.Name, "description": dealA.Description})
	if result.Error != nil {

		r.JSON(400, gin.H{
			"msg": "invalid json",
		})
		r.Abort()

		return
	}
	r.JSON(200, deal)
}

func DeleteDeal(r *gin.Context) {
	var deal Deal
	id := r.Param("id")
	result := database.GlobalDB.Delete(&deal, "id = ?", id)
	if result.Error != nil {
		r.JSON(400, gin.H{
			"msg": "invalid json",
		})
		r.Abort()

		return
	}
	r.JSON(200, true)
}
