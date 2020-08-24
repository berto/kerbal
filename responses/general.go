package responses

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

func newResponse(status int, body string) events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{
		StatusCode: status,
		Headers: map[string]string{
			"Content-Type":                     "application/json",
			"Access-Control-Allow-Origin":      "*",
			"Access-Control-Allow-Credentials": "true",
			"Access-Control-Allow-Headers":     "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With",
			"Access-Control-Allow-Methods":     "POST, OPTIONS, GET, PUT",
		},
		Body: body,
	}
}

// ClientError sends client error
func ClientError(err error) events.APIGatewayProxyResponse {
	return newResponse(http.StatusBadRequest, fmt.Sprintf(`{"error":"%s"}`, err.Error()))
}

// ServerError sends server error
func ServerError(err error) events.APIGatewayProxyResponse {
	return newResponse(http.StatusInternalServerError, fmt.Sprintf(`{"error":"%s"}`, err.Error()))
}

// OK sends data
func OK(data interface{}) events.APIGatewayProxyResponse {
	body, err := json.Marshal(data)
	if err != nil {
		return ServerError(err)
	}
	return newResponse(http.StatusOK, string(body))
}
