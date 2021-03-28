package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"
	"url-shortner/src/dependency"
	"url-shortner/util"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/parser"
	"github.com/guregu/dynamo"
)

var BaseResourcePath string = ""

// Models our json input
type ServerEnv struct {
	DynamoRegion     string `json:"dynamoRegion"`
	IsOffline        bool   `json:"isOffline"`
	BaseResourcePath string `json:"baseResourcePath"`
}

// reads the enviroment file if it exists , if not then there's a problem
func readEnviromentFile() *ServerEnv {
	var serverEnv ServerEnv
	jsonFile, err := os.Open("dist/env.json")
	if err != nil {
		fmt.Printf(":0 PANIK , enviroment json not attempting local dir found")
		jsonFile, err = os.Open("env.json")
		if err != nil {
			panic(":0 PANIK , enviroment json not found")
		}
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	err = json.Unmarshal(byteValue, &serverEnv)
	if err != nil {
		panic(":0 PANIK , enviroment json could not decode")
	}
	return &serverEnv
}

// fetch the db connection here
func getDB(serverEnv *ServerEnv) *dynamo.DB {
	return dynamo.New(session.New(), &aws.Config{Region: aws.String(serverEnv.DynamoRegion)})
}

func makeCors(router *gin.Engine) {
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://shrter.link", "https://shrter.xyz", "http://127.0.0.1:5500", "http://localhost:8080"},
		AllowMethods:     []string{"PUT", "PATCH", "POST", "DELETE", "GET"},
		AllowHeaders:     []string{"Origin", "content-type", "authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
}

func makeStatic(router *gin.Engine) {
	router.LoadHTMLGlob(BaseResourcePath + "client/*.html")
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
	serverEnv := readEnviromentFile()
	BaseResourcePath = serverEnv.BaseResourcePath
	db := getDB(serverEnv)
	controller := dependency.MakeDependencies(db)
	router := gin.Default()
	makeRouter(router, controller)
	if serverEnv.IsOffline {
		_ = router.Run(":7989")

	} else {
		// handle as lambda enviroment
		lambda.Start(util.Handler(router))
	}

}
