package handler

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
)

type IHandler interface {
	Handle(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error)
}
