package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func createItem(c *gin.Context, Db *gorm.DB) {

	var item Item
	if err := c.ShouldBindJSON(&item); err != nil {
		log.Fatal("There's was an error!, ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	Db.Create(&item)
	c.JSON(http.StatusCreated, item)
}

func getItem(c *gin.Context, Db *gorm.DB) {
	var itemz []Item
	if err := Db.Find(&itemz).Error; err != nil {
		log.Fatal("There's was an error!, ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't find anything"})
		return
	}
	c.JSON(http.StatusOK, itemz)
}

func updateItem(c *gin.Context, Db *gorm.DB) {
	id := c.Param("id")
	var item Item
	if err := Db.First(&item, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Предмет не найден"})
		return
	}
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверные данные"})
		return
	}
	Db.Save(&item)
	c.JSON(http.StatusOK, item)
}

func deleteItem(c *gin.Context, Db *gorm.DB) {
	id := c.Param("id")
	if err := Db.Delete(&Item{}, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Предмет не найден"})
		return
	}
	c.JSON(http.StatusNoContent, nil) // Возвращаем пустой ответ
}
