// TODO: setup https://github.com/mefellows/vesper
package middleware

import (
	database "slanj/common/database/repository"

	"github.com/aws/aws-lambda-go/events"
)

var repository = &database.UserRepository{}

func TokenMiddleware(request events.APIGatewayProxyRequest) string {
	token := request.Headers["Authorization"]
	repository.SaveUser(token)

	return token
}
