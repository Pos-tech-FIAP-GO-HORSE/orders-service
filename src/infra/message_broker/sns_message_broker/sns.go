package sns_message_broker

import (
	"context"

	"github.com/Pos-tech-FIAP-GO-HORSE/orders-service/src/infra/message_broker"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sns"
)

type SNSMessageBroker struct {
	client *sns.Client
}

func NewSNSMessageBroker(client *sns.Client) message_broker.IMessageBroker {
	return &SNSMessageBroker{client}
}

func (ref *SNSMessageBroker) Publish(ctx context.Context, topic, message string) error {
	publishInput := sns.PublishInput{
		TopicArn: aws.String(topic),
		Message:  aws.String(message),
	}

	_, err := ref.client.Publish(ctx, &publishInput)
	if err != nil {
		return err
	}

	return nil
}

func (ref *SNSMessageBroker) Subscribe(ctx context.Context, topic, protocol string) (string, error) {
	subscribeInput := &sns.SubscribeInput{
		TopicArn: aws.String(topic),
		Protocol: aws.String(protocol),
	}

	result, err := ref.client.Subscribe(ctx, subscribeInput)
	if err != nil {
		return "", err
	}

	return *result.SubscriptionArn, nil
}
