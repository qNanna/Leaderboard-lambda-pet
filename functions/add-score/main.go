package main

import (
	config "slanj/common/config"
	database "slanj/common/database"
	handler "slanj/functions/add-score/presentation"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	config.Init()
	database.Init()
	lambda.Start(handler.AddScoreHandler)
}
