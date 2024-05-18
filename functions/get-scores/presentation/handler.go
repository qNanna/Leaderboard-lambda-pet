package presentation

import (
	"context"
	"encoding/json"

	application "slanj/functions/get-scores/application"

	"github.com/aws/aws-lambda-go/events"
)

var service = &application.Service{}

func ScoreHandler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	data := service.GetScores()

	response, _ := json.Marshal(&data)

	return events.APIGatewayProxyResponse{
		Body:       string(response),
		StatusCode: 200,
		Headers:    map[string]string{"Content-Type": "application/json"},
	}, nil

}
