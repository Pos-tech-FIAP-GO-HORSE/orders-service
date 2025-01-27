package main

import (
	"context"
	"os"
	"time"

	"github.com/Pos-tech-FIAP-GO-HORSE/orders-service/src/core/service/order_service"
	"github.com/Pos-tech-FIAP-GO-HORSE/orders-service/src/function/update_order_status/handler"
	"github.com/Pos-tech-FIAP-GO-HORSE/orders-service/src/infra/message_broker/sns_message_broker"
	"github.com/Pos-tech-FIAP-GO-HORSE/orders-service/src/infra/repository/mongodb_repository"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sns"
	_ "github.com/joho/godotenv/autoload"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

func init() {
	if os.Getenv("APP_ENV") == "production" {
		zap.ReplaceGlobals(zap.Must(zap.NewProduction()))
	} else {
		zap.ReplaceGlobals(zap.Must(zap.NewDevelopment()))
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	clientOptions := options.Client().ApplyURI(os.Getenv("DB_URI"))
	dbClient, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		zap.L().Fatal("unable to connect on database", zap.Error(err))
	}

	zap.L().Info("database connected successfully")

	database := dbClient.Database(os.Getenv("DB_NAME"))
	ordersCollection := database.Collection("orders")

	cfg, err := config.LoadDefaultConfig(
		ctx,
		config.WithRegion("us-east-1"),
		config.WithBaseEndpoint(os.Getenv("SNS_URL")),
	)
	if err != nil {
		zap.L().Fatal("unable to load aws config", zap.Error(err))
	}

	topics := map[string]string{
		"order-updated": os.Getenv("TOPIC_ORDER_UPDATED"),
	}

	snsClient := sns_message_broker.NewSNSMessageBroker(sns.NewFromConfig(cfg))

	orderRepository := mongodb_repository.NewOrderRepository(ordersCollection)
	orderService := order_service.NewOrderService(orderRepository, snsClient, topics)
	handler := handler.NewHandler(orderService)

	lambda.Start(handler.Handle)
}
