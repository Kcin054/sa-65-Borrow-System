package controller

import (
	"github.com/PatiphatPanchad/sa-65-example/entity"

	"github.com/gin-gonic/gin"

	"net/http"
)

// POST /List_datas

func CreateList_data(c *gin.Context) {

	var list_data entity.List_data
	var equipment entity.Equipment
	var borrow_card entity.Borrow_card
	var room entity.Room

	if err := c.ShouldBindJSON(&list_data); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	if tx := entity.DB().Where("id = ?", list_data.Equipment_id).First(&equipment); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "equipment หาไม่เจอเวนนนน"})
		return

	}
	if tx := entity.DB().Where("id = ?", list_data.Borrow_card_id).First(&borrow_card); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "room price หาไม่เจอเวนนนน"})
		return

	}
	if tx := entity.DB().Where("id = ?", list_data.Room_id).First(&room); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "set of furniture หาไม่เจอเวนนนน"})

		return

	}
	wv := entity.List_data{
		Equipment:   equipment,
		Borrow_card: borrow_card,
		Room:        room,
		Borrow_time: list_data.Borrow_time,
		Return_time: list_data.Return_time,
	}
	if err := entity.DB().Create(&wv).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": wv})
}

// GET /List_data/:id
func GetList_data(c *gin.Context) {
	var List_data entity.List_data
	id := c.Param("id")
	if tx := entity.DB().Where("id = ?", id).First(&List_data); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "List_data not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": List_data})
}

// GET /list_datas
func ListList_data(c *gin.Context) {
	var List_data []entity.List_data
	if err := entity.DB().Preload("Equipment").Preload("Borrow_card").Preload("Room").Raw("SELECT * FROM List_data").Find(&List_data).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": List_data})
}

// DELETE /list_datas/:id
func DeleteList_data(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM List_datas WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "List_data not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /list_datas
func UpdateList_data(c *gin.Context) {
	var List_data entity.List_data
	if err := c.ShouldBindJSON(&List_data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", List_data.ID).First(&List_data); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "List_data not found"})
		return
	}

	if err := entity.DB().Save(&List_data).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": List_data})
}
