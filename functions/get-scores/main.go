package main

import (
	config "slanj/common/config"
	database "slanj/common/database"
	handler "slanj/functions/get-scores/presentation"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	config.Init()
	database.Init()
	lambda.Start(handler.ScoreHandler)
}
