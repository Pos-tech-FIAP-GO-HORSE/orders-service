package main

import (
	"context"
	"os"
	"time"

	"github.com/Pos-tech-FIAP-GO-HORSE/orders-service/src/core/service/order_service"
	"github.com/Pos-tech-FIAP-GO-HORSE/orders-service/src/infra/repository/mongodb_repository"
	"github.com/aws/aws-lambda-go/lambda"
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
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		zap.L().Fatal("unable to connect on database", zap.Error(err))
	}

	database := client.Database(os.Getenv("DB_NAME"))
	ordersCollection := database.Collection("orders")

	orderRepository := mongodb_repository.NewOrderRepository(ordersCollection)
	orderService := order_service.NewOrderService(orderRepository)
	handler := NewHandler(orderService)

	lambda.Start(handler.Handle)
}
