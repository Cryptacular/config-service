package main

import (
	"bytes"
	"encoding/xml"
	"errors"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"log"
)

var (
	errNameNotProvided = errors.New("no name was provided in the HTTP body")
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Printf("Processing Lambda request %s\n", request.RequestContext.RequestID)

	x := createXMLString()

	return events.APIGatewayProxyResponse{
		Body:       x,
		StatusCode: 200,
		Headers:    map[string]string{"Accept": "application/xml"},
	}, nil
}

func main() {
	lambda.Start(handler)
}

func createXMLString() string {
	data := config{}

	buf := new(bytes.Buffer)
	enc := xml.NewEncoder(buf)
	enc.Indent("", "    ")
	enc.Encode(data)

	return buf.String()
}
