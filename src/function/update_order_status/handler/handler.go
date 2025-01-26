package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/Pos-tech-FIAP-GO-HORSE/orders-service/src/core/service/order_service"
	"github.com/Pos-tech-FIAP-GO-HORSE/orders-service/src/function"
	"github.com/Pos-tech-FIAP-GO-HORSE/orders-service/src/function/update_order_status/contract"
	"github.com/aws/aws-lambda-go/events"
)

type Handler struct {
	orderService order_service.IOrderService
}

func NewHandler(orderService order_service.IOrderService) function.IHandler {
	return &Handler{orderService}
}

func (ref *Handler) Handle(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	orderID, ok := req.PathParameters["id"]
	if !ok {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Headers: map[string]string{
				"Content-Type": "application-json",
			},
			Body: "id path parameter missing",
		}, nil
	}

	var orderRequest contract.UpdateOrderRequest

	if err := json.Unmarshal([]byte(req.Body), &orderRequest); err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Headers: map[string]string{
				"Content-Type": "application-json",
			},
			Body: err.Error(),
		}, nil
	}

	updatedOrder, err := ref.orderService.UpdateStatusByID(ctx, orderID, orderRequest.Status)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError, // TODO: alterar para pegar o status code dinamicamente conforme o retorno da camada de negócios
			Headers: map[string]string{
				"Content-Type": "application-json",
			},
			Body: err.Error(),
		}, nil
	}

	response := contract.UpdateOrderResponseFromDomain(updatedOrder)
	rawResponse, _ := json.Marshal(response)

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusCreated,
		Headers: map[string]string{
			"Content-Type": "application-json",
		},
		Body: string(rawResponse),
	}, nil
}
