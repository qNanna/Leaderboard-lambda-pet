package main

import (
	config "slanj/src/common/config"
	database "slanj/src/common/database"
	handler "slanj/src/functions/get-scores/presentation"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	config.Init()
	database.Init()
	lambda.Start(handler.ScoreHandler)
}
