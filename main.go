package main

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/Cryptacular/config-service/schema"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/neelance/graphql-go"
	"log"
)

var (
	// ErrQueryNameNotProvided is returned when no query was provided in the HTTP body
	ErrQueryNameNotProvided = errors.New("no query was provided in the HTTP body")
	gqlSchema               *graphql.Schema
)

func init() {
	gqlSchema = graphql.MustParseSchema(schema.Schema, &schema.Resolver{})
}

func main() {
	println("Starting Config Service...")
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

	response := gqlSchema.Exec(context, params.Query, params.OperationName, params.Variables)
	responseJSON, err := json.Marshal(response)

	if err != nil {
		log.Print("Could not serialise body")
	}

	return events.APIGatewayProxyResponse{
		Body:       string(responseJSON),
		StatusCode: 200,
	}, nil
}
