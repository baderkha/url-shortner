package util

import (
	"log"

	"github.com/aws/aws-lambda-go/events"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
)

var initialized = false
var ginLambda *ginadapter.GinLambda

// Handler - lambda handler proxy
func Handler(r *gin.Engine) func(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return func(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		// stdout and stderr are sent to AWS CloudWatch Logs
		log.Printf("Recieved: %s", req.Body)

		if !initialized {
			ginLambda = ginadapter.New(r)
			initialized = true
		}

		// If no name is provided in the HTTP request body, throw an error
		return ginLambda.Proxy(req)
	}
}
