package main_test

import (
	"context"
	"testing"

	main "github.com/Cryptacular/config-service"

	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/assert"
)

func TestHandler(t *testing.T) {
	ctx := context.TODO()
	tests := []struct {
		request events.APIGatewayProxyRequest
		expect  string
		context context.Context
		err     error
	}{
		{
			// Test that the handler responds with the correct response
			// when a valid name is provided in the HTTP body
			request: events.APIGatewayProxyRequest{
				Body: `{"query":"query test {\n configuration(id:\"1\") {\n id\n country\n }\n abTests(id: \"1\") {\n id\n userId\n name\n segment\n expiry\n}\n}\n","variables":null,"operationName":"test"}`,
			},
			expect:  "{\"data\":{\"configuration\":{\"id\":\"1\",\"country\":\"NZ\"},\"abTests\":[{\"id\":\"1\",\"userId\":\"1\",\"name\":\"NZ\",\"segment\":\"com\",\"expiry\":\"2018-12-31 00:00:00 +0000 UTC\"}]}}",
			context: ctx,
			err:     nil,
		},
	}

	for _, test := range tests {
		response, err := main.Handler(test.context, test.request)
		assert.IsType(t, test.err, err)
		assert.Equal(t, test.expect, response.Body)
	}
}
