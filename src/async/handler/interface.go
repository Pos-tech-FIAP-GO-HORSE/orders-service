package handler

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
)

type IHandler interface {
	Handle(ctx context.Context, snsEvent events.SNSEvent) error
}
