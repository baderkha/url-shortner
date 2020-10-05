package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/parser"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"io/ioutil"
	"log"
	"os"
	"url-shortner/src/dependency"
	"url-shortner/src/repository"
)

// Models our json input
type ServerEnv struct {
	DSN string `json:"dsn"`
}

// reads the enviroment file if it exists , if not then there's a problem
func readEnviromentFile() *ServerEnv {
	var serverEnv ServerEnv
	jsonFile, err := os.Open("env.json")
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
	db.Logger = logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			LogLevel: logger.Info, // Log level
			Colorful: true,        // Disable color
		},
	)
	return db
}

func migrate(db *gorm.DB) {
	var link repository.Link
	_ = db.AutoMigrate(&link)
}

// Define the gin routes in here using the router
func makeRoutes(router *gin.Engine, controller *dependency.Dependency) {
	router.GET("", func(c *gin.Context) {
		extensions := parser.CommonExtensions | parser.AutoHeadingIDs
		parser := parser.NewWithExtensions(extensions)
		file, err := os.Open("README.md")
		if err != nil {
			panic(":0 PANIK , enviroment json not found")
		}
		byteValue, _ := ioutil.ReadAll(file)
		html := markdown.ToHTML(byteValue, parser, nil)
		c.Data(200, "text/html; charset=utf-8", html)
	})
	router.GET("links/:id", controller.FetchLink)
	router.POST("links", controller.ShortenLink)
}

// no buisness logic lives here
func main() {
	db := getDB(readEnviromentFile())
	migrate(db)
	controller := dependency.MakeDependencies(db)
	router := gin.Default()
	makeRoutes(router, controller)
	_ = router.Run()
}
