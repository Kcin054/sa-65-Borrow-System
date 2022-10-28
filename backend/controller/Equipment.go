package controller

import (
	"github.com/PatiphatPanchad/sa-65-example/entity"

	"github.com/gin-gonic/gin"

	"net/http"
)

// POST /Equipments

func CreateEquipment(c *gin.Context) {

	var equipments entity.Equipment

	if err := c.ShouldBindJSON(&equipments); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	if err := entity.DB().Create(&equipments).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": equipments})

}

// GET /Equipment/:id

func GetEquipment(c *gin.Context) {

	var equipments entity.Equipment

	id := c.Param("id")

	if err := entity.DB().Raw("SELECT * FROM Equipment WHERE id = ?", id).Scan(&equipments).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": equipments})

}

// GET /Equipments

func ListEquipments(c *gin.Context) {

	var equipments []entity.Equipment

	if err := entity.DB().Raw("SELECT * FROM Equipment").Scan(&equipments).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": equipments})

}

// DELETE /Equipments/:id

func DeleteEquipment(c *gin.Context) {

	id := c.Param("id")

	if tx := entity.DB().Exec("DELETE FROM Equipment WHERE id = ?", id); tx.RowsAffected == 0 {

		c.JSON(http.StatusBadRequest, gin.H{"error": "Equipment not found"})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": id})

}

// PATCH /Equipments

func UpdateEquipment(c *gin.Context) {

	var equipments entity.Equipment

	if err := c.ShouldBindJSON(&equipments); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	if tx := entity.DB().Where("id = ?", equipments.ID).First(&equipments); tx.RowsAffected == 0 {

		c.JSON(http.StatusBadRequest, gin.H{"error": "Equipment not found"})

		return

	}

	if err := entity.DB().Save(&equipments).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": equipments})

}
