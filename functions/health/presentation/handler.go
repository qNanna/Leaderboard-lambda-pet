package presentation

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
)

func HealthHandler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{Body: "Working 📶", StatusCode: 200, Headers: map[string]string{"Content-Type": "application/json"}}, nil
}
