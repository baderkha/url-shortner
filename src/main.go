package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"io/ioutil"
	"os"
	"url-shortner/src/dependency"
)

// Models our json input
type ServerEnv struct {
	DSN string `json:"dsn"`
}

// reads the enviroment file if it exists , if not then there's a problem
func readEnviromentFile() *ServerEnv {
	var serverEnv ServerEnv
	jsonFile, err := os.Open("../env.json")
	if err != nil {
		panic(":0 PANIK , enviroment json not found")
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	err = json.Unmarshal(byteValue, &serverEnv)
	if err != nil {
		panic(":0 PANIK , enviroment json could not decode")
	}
	return &serverEnv
}

// fetch the db connection here
func getDB(serverEnv *ServerEnv) *gorm.DB {
	db, err := gorm.Open(mysql.Open(serverEnv.DSN), &gorm.Config{})
	if err != nil {
		panic(":0 PANIk , DB CONN COULD NOT BE MADE")
	}
	return db
}

// Define the gin routes in here using the router
func makeRoutes(router *gin.Engine, controller *dependency.Dependency) {
	router.GET("links/:id", controller.FetchLink)
	router.POST("links", controller.ShortenLink)
}

// no buisness logic lives here
func main() {
	db := getDB(readEnviromentFile())
	controller := dependency.MakeDependencies(db)
	router := gin.Default()
	makeRoutes(router, controller)
	_ = router.Run()
}
