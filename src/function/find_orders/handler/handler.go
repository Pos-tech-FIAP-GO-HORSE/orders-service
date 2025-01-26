package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/Pos-tech-FIAP-GO-HORSE/orders-service/src/core/service/order_service"
	"github.com/Pos-tech-FIAP-GO-HORSE/orders-service/src/function"
	"github.com/Pos-tech-FIAP-GO-HORSE/orders-service/src/function/find_orders/contract"
	"github.com/aws/aws-lambda-go/events"
)

type Handler struct {
	orderService order_service.IOrderService
}

func NewHandler(orderService order_service.IOrderService) function.IHandler {
	return &Handler{orderService}
}

func (ref *Handler) Handle(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	orders, err := ref.orderService.Find(ctx)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError, // TODO: alterar para pegar o status code dinamicamente conforme o retorno da camada de negócios
			Headers: map[string]string{
				"Content-Type": "application-json",
			},
			Body: err.Error(),
		}, nil
	}

	response := contract.OrdersFromDomain(orders)
	rawResponse, _ := json.Marshal(response)

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Headers: map[string]string{
			"Content-Type": "application-json",
		},
		Body: string(rawResponse),
	}, nil
}
