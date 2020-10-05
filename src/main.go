package main

import (
	"encoding/json"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/parser"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"io/ioutil"
	"log"
	"os"
	"time"
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

func makeCors(router *gin.Engine) {
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://shrter.xyz", "http://127.0.0.1:5500", "http://localhost:8080"},
		AllowMethods:     []string{"PUT", "PATCH", "POST", "DELETE", "GET"},
		AllowHeaders:     []string{"Origin", "content-type", "authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
}

func makeStatic(router *gin.Engine) {
	router.LoadHTMLGlob("client/*.html")
	router.GET("", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})
	router.GET("/readme", func(c *gin.Context) {
		extensions := parser.CommonExtensions | parser.AutoHeadingIDs
		parser := parser.NewWithExtensions(extensions)
		file, _ := os.Open("README.md")
		byteValue, _ := ioutil.ReadAll(file)
		html := markdown.ToHTML(byteValue, parser, nil)
		c.Data(200, "text/html; charset=utf-8", html)
	})
}

// Define the gin routes in here using the router
func makeRestRoutes(router *gin.Engine, controller *dependency.Dependency) {
	api := router.Group("api")
	{
		// link forward , hitting this route will cause the link to be resolved and forwarded in the browser
		api.GET("forward/links/:id", controller.ForwardLink)
		// get them the json record
		api.GET("links/:id", controller.FetchLink)
		// generate a new link
		api.POST("links", controller.ShortenLink)
	}
}

func makeRouter(router *gin.Engine, controller *dependency.Dependency) {
	makeCors(router)
	makeStatic(router)
	makeRestRoutes(router, controller)
}

// no buisness logic lives here
func main() {
	db := getDB(readEnviromentFile())
	migrate(db)
	controller := dependency.MakeDependencies(db)
	router := gin.Default()
	makeRouter(router, controller)
	_ = router.Run()
}
