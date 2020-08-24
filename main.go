package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/berto/kerbal/controllers"
	"github.com/berto/kerbal/responses"
	"github.com/pkg/errors"
)

func handleLambda(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	switch request.HTTPMethod {
	case http.MethodOptions:
		return responses.OK("ok"), nil
	case http.MethodGet:
		items, err := controllers.GetItems(ctx)
		if err != nil {
			return responses.ServerError(err), nil
		}
		return responses.OK(items), nil
	case http.MethodPost:
		input := controllers.KerbalItems{}
		if err := json.Unmarshal([]byte(request.Body), &input); err != nil {
			return responses.ClientError(errors.Wrap(err, request.Body)), nil
		}
		if err := input.Validate(); err != nil {
			return responses.ClientError(err), nil
		}
		id, err := controllers.CreateKerbal(ctx, input)
		if err != nil {
			return responses.ServerError(err), nil
		}
		return responses.OK(map[string]string{"id": id}), nil
	default:
		return responses.ClientError(errors.New("invalid method")), nil
	}
}

func main() {
	lambda.Start(handleLambda)
}
