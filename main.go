package main

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Item struct {
	Id          int       `gorm:"primaryKey;autoIncrement"`
	Name        string    `gorm:"not null"`
	Description string    `gorm:"default:''"`
	Price       float64   `gorm:"not null"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
}

func main() {
	dsn := "host=127.0.0.1 user=postgres password=Maks2005 dbname=Task port=5433 sslmode=disable"

	DataBase, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("There's error with connection: ", err)
	}

	DataBase.AutoMigrate(&Item{})
	r := gin.Default()
	r.POST("/items", func(c *gin.Context) {
		createItem(c, DataBase)
	})

	r.GET("/items", func(c *gin.Context) {
		getItem(c, DataBase)
	})

	r.PUT("/items", func(c *gin.Context) {
		updateItem(c, DataBase)
	})

	r.DELETE("/items", func(c *gin.Context) {
		deleteItem(c, DataBase)
	})

	r.Run(":8080")
}
