package main

import (
	"ap/docs"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// 佈告欄的資訊列表
var data []*BulletinResponse

// @version 1.0
// @description This is a user API for a bulletin board system.

func main() {
	r := gin.Default()
	r.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, PUT, POST, DELETE",
		RequestHeaders:  "Origin, Authorization, Content-Type",
		ExposedHeaders:  "",
		MaxAge:          50 * time.Second,
		ValidateHeaders: false,
	}))
	// 將api文件公開到指定路徑
	docs.Init()
	r.Use(static.Serve("/", static.LocalFile("./dist", true)))
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// 註冊路由
	r.POST("/announcements", AddBulletin)          // 新增佈告
	r.DELETE("/announcements/:id", DeleteBulletin) // 刪除佈告
	r.PUT("/announcements/:id", WriteBulletin)     // 編輯佈告
	r.GET("/announcements", BulletinList)          // 取得布告

	// 開啟服務
	r.Run(":3000")
}

// AddBulletinRequest request body for add BulletinResponse
// @type body
// @param title body string true "title"
// @param content body string true "content"
// @success 200 {object} BulletinResponse
type AddBulletinRequest struct {
	Title   string `form:"title" json:"title" binding:"required"`
	Content string `form:"content" json:"content" binding:"required"`
}

// BulletinResponse response body for add BulletinResponse
// @type object
// @property id integer
// @property title string
// @property content string
type BulletinResponse struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

// AddBulletin godoc
// @Summary 新增佈告
// @Description 發送主題和內容 新增佈告欄公告
// @Tags user
// @Accept  json
// @Produce  json
// @Param body body AddBulletinRequest true "AddBulletin Request"
// @Success 200 {object} BulletinResponse
// @Router /announcements [post]
func AddBulletin(c *gin.Context) {
	// Parse request parameters
	var request AddBulletinRequest
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	BulletinID := 1
	// 目前最後一筆的編號加一
	if len(data) == 0 {
		BulletinID = 1
	} else {
		BulletinID = data[len(data)-1].ID + 1
	}
	response := &BulletinResponse{
		ID:      BulletinID,
		Title:   request.Title,
		Content: request.Content,
	}

	data = append(data, response)
	c.JSON(http.StatusOK, data)
}

// DeleteBulletinRequest request body for DeleteResponse
// @type body
// @param id body int true "id"
// @success 200 {object} DeleteBulletinResponse
type DeleteBulletinRequest struct {
	ID int `form:"id" json:"id" binding:"required"` // 佈告欄編號
}

// DeleteBulletinResponse response body for DeleteResponse
// @type object
// @property id integer
// @property title string
// @property content string
type DeleteBulletinResponse struct {
	ID      int    `json:"id"`      // 佈告欄編號
	Title   string `json:"title"`   // 佈告欄主旨
	Content string `json:"content"` // 佈告欄內文
}

// @Summary 刪除公告
// @Description 發送編號 刪除佈告欄公告
// @ID announcements-delete
// @Tags user
// @Accept      json
// @Produce     json
// @Param id path int true "佈告ID"
// @Success 200 {Success} string "刪除成功"
// @Failure 400 {string} string "輸入不正確"
// @Failure 404 {string} string "找不到資源"
// @Router /announcements/{id} [delete]
func DeleteBulletin(c *gin.Context) {

	var id = c.Param("id")

	BulletinListID, error := strconv.Atoi(id)
	if error != nil {
		log.Print(BulletinListID, error)
		c.JSON(http.StatusOK, error)
	}

	for i, element := range data {
		//索引值 i 的 編號值若等於想刪除的內容
		//新的數據將等於 陣列內的索引[1,2,3,4,5,6]  [5] [1,2,3,4,6]
		if element.ID == BulletinListID {
			data = append(data[:i], data[i+1:]...)
		}

	}
	var keyfind []int
	for i, _ := range data {
		keyfind = append(keyfind, i)
	}

	c.JSON(http.StatusOK, "Success")
	return

}

// WriteBulletinRequest request body for WriteResponse
// @type body
// @param id body int true "id"
// @param title body string true "title"
// @param content body string true "content"
// @success 200 {object} WriteBulletinResponse
type WriteBulletinRequest struct {
	ID      int    `form:"id" json:"id" binding:"required"`
	Title   string `form:"title" json:"title" binding:"required"`
	Content string `form:"content" json:"content" binding:"required"`
}

// WriteBulletinResponse response body for WriteResponse
// @type object
// @property id integer
// @property title string
// @property content string
type WriteBulletinResponse struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

// WriteBulletin godoc
// @Summary 編輯佈告
// @Description 發送主題和內容 新增佈告欄公告
// @Tags user
// @Accept  json
// @Produce  json
// @Param body body WriteBulletinRequest true "WriteBulletin Request"
// @Success 200 {object} WriteBulletinResponse
// @Router /announcements/:id [put]
func WriteBulletin(c *gin.Context) {
	// Parse request parameters
	var request WriteBulletinRequest
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for _, element := range data {
		if element.ID == request.ID {
			element.Title = request.Title
			element.Content = request.Content
		}
	}
	c.JSON(http.StatusOK, data)
	return
}

// BulletinListRequest request body for BulletinListRequest
// @type body
// @success 200 {object} BulletinListResponse
type BulletinListRequest struct {
}

// BulletinListResponse response body for add BulletinResponse
// @type object
// @property id integer
// @property title string
// @property content string
type BulletinListResponse struct {
	ID      int    `json:"id"`      // 佈告欄編號
	Title   string `json:"title"`   // 佈告欄主旨
	Content string `json:"content"` // 佈告欄內文
}

// BulletinList godoc
// @Summary 取得佈告列表
// @Description 佈告列表路徑 提供當前所有佈告內容
// @Tags user
// @Accept  json
// @Produce  json
// @Param body body BulletinListRequest true "BulletinList Request"
// @Success 200 {object} BulletinListResponse
// @Router /announcements [get]
func BulletinList(c *gin.Context) {
	// Parse request parameters
	var request BulletinListRequest
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, data)
}
