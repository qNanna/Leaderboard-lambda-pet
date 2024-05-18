package main

import (
	handler "slanj/functions/health/presentation"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(handler.HealthHandler)
}
