package imagechild

import (
	"Go-market-test/pkg/database"
	deal2 "Go-market-test/pkg/deal"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"io"
	"log"
	"net/http"
	"os"
	filepath2 "path/filepath"
	"strconv"
	"strings"
)

type ImageChild struct {
	gorm.Model
	Xid       string
	FileName  string
	Extension string
	ParentId  uint
}

func GetImageChild(id uint) (ic ImageChild, err error) {
	result := database.GlobalDB.Where("id=?", id).First(&ic)
	if result.Error != nil {
		return ic, err
	}
	return ic, err
}

func GetImageChildByXid(xid string) (ic ImageChild, err error) {
	result := database.GlobalDB.Where("xid=?", xid).First(&ic)
	if result.Error != nil {
		return ic, err
	}
	return ic, err
}

func GetDealImageChilds(dealId uint) (ics ImageChild, err error) {
	result := database.GlobalDB.Where("parent_id = ?").Find(&ics)
	if result.Error != nil {
		return ics, err
	}
	return ics, err
}

func Upload(c *gin.Context) {
	id := c.Param("id")
	var deal deal2.Deal
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("file err : %s", err.Error()))
		return
	}
	var ext = filepath2.Ext(header.Filename)
	g := uuid.New()
	filename := g.String() + ext
	out, err := os.Create("data/" + filename)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		log.Fatal(err)
	}

	var ic ImageChild
	ic.Xid = filename
	ic.FileName = header.Filename
	ic.Extension = ext

	parentId, ok := c.Get("parentId")
	if ok {
		pi, err := strconv.Atoi(parentId.(string))
		if err != nil {
			return
		}
		ic.ParentId = uint(pi)
	}

	_ = database.GlobalDB.First(&deal, id)

	ic.ParentId = deal.ID

	_ = database.GlobalDB.Create(&ic)

	var abs string
	//abs, err = filepath2.Abs("data/" + ic.Xid)
	//if err != nil {
	//	c.JSON(404, "")
	//}
	abs = "http://127.0.0.1:8087/data/" + ic.Xid
	abs = strings.Replace(abs, "\\", "/", -1)
	deal.Image = abs
	database.GlobalDB.Model(deal).Update("image", deal.Image)

	c.JSON(http.StatusOK, gin.H{"filepath": filename})
}

func UserFileDownloadCommonService(c *gin.Context) {
	var err error
	var ic ImageChild
	ic, err = GetImageChildByXid(c.Param("xid"))
	if err != nil {
		c.JSON(404, "")
	}

	if ic.ID != 0 {
		var abs string
		abs, err = filepath2.Abs("data/" + ic.Xid)
		if err != nil {
			c.JSON(404, "")
		}
		c.JSON(http.StatusOK, abs)
	}
	c.JSON(404, "")
}
