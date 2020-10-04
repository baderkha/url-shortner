package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"url-shortner/dependency"
)

func main() {
	dsn := "root:password@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(":0 PANIk")
	}
	controller := dependency.MakeDependencies(db)
	router := gin.Default()
	router.GET("links/:id", controller.FetchLink)
	router.POST("links", controller.ShortenLink)
	router.Run()
}
