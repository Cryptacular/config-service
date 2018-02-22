package main

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/Cryptacular/config-service/database"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/neelance/graphql-go"
	"log"
)

var (
	// ErrQueryNameNotProvided is returned when no query was provided in the HTTP body
	ErrQueryNameNotProvided = errors.New("no query was provided in the HTTP body")
	schema                  *graphql.Schema
)

func init() {
	schema = graphql.MustParseSchema(database.Schema, &database.Resolver{})
}

func main() {
	lambda.Start(Handler)
}

// Handler for AWS Lambda
func Handler(context context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Printf("Processing Lambda request %s\n", request.RequestContext.RequestID)

	if len(request.Body) < 1 {
		return events.APIGatewayProxyResponse{}, ErrQueryNameNotProvided
	}

	var params struct {
		Query         string                 `json:"query"`
		OperationName string                 `json:"operationName"`
		Variables     map[string]interface{} `json:"variables"`
	}

	if err := json.Unmarshal([]byte(request.Body), &params); err != nil {
		log.Print("Could not deserialise body", err)
	}

	response := schema.Exec(context, params.Query, params.OperationName, params.Variables)
	responseJSON, err := json.Marshal(response)

	if err != nil {
		log.Print("Could not serialise body")
	}

	return events.APIGatewayProxyResponse{
		Body:       string(responseJSON),
		StatusCode: 200,
	}, nil
}
