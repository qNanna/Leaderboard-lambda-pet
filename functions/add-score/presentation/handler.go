package presentation

import (
	"context"
	"encoding/json"

	application "slanj/functions/add-score/application"
	middleware "slanj/middlewares"

	"github.com/aws/aws-lambda-go/events"
)

var service = &application.Service{}

type RequestBody struct {
	Score int `json:"score"`
}

func AddScoreHandler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	token := middleware.TokenMiddleware(request)

	var data RequestBody
	json.Unmarshal([]byte(request.Body), &data)

	response := service.SaveScore(token, data.Score)
	if !response {
		panic("Unsuccess")
	}

	return events.APIGatewayProxyResponse{Body: "Done", StatusCode: 200, Headers: map[string]string{"Content-Type": "application/json"}}, nil
}

// 76561198033666298
// 76561198059258743
