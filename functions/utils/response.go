package utils

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
)

func APIGatewayProxyResponseWithJSON(code int, v any) (events.APIGatewayProxyResponse, error) {
	bodyBytes, err := json.Marshal(v)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	return events.APIGatewayProxyResponse{
		StatusCode: code,
		Headers:    map[string]string{"Content-Type": "application/json"},
		Body:       string(bodyBytes),
	}, nil
}
