package message_broker

import "context"

type IMessageBroker interface {
	Publish(ctx context.Context, topicArn, message string) error
	Subscribe(ctx context.Context, topicArn, protocol string) (string, error)
}
