package handler

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/Pos-tech-FIAP-GO-HORSE/orders-service/src/core/service/order_service"
	"github.com/aws/aws-lambda-go/events"
)

type Handler struct {
	orderService order_service.IOrderService
}

func NewHandler(orderService order_service.IOrderService) IHandler {
	return &Handler{orderService}
}

func (ref *Handler) Handle(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	switch req.HTTPMethod {
	case http.MethodPost:
		responseRaw, statusCode, err := ref.create(ctx, []byte(req.Body))
		response := string(responseRaw)
		if err != nil {
			response = err.Error()
		}

		return events.APIGatewayProxyResponse{
			StatusCode: statusCode,
			Headers: map[string]string{
				"Content-Type": "application-json",
			},
			Body: response,
		}, nil

	case http.MethodGet:
		var (
			responseRaw []byte
			statusCode  int
			err         error
		)

		id, exists := req.PathParameters["id"]
		if exists {
			responseRaw, statusCode, err = ref.findByID(ctx, id)

		} else {
			responseRaw, statusCode, err = ref.find(ctx)
		}

		response := string(responseRaw)

		if err != nil {
			response = err.Error()
		}

		return events.APIGatewayProxyResponse{
			StatusCode: statusCode,
			Headers: map[string]string{
				"Content-Type": "application-json",
			},
			Body: response,
		}, nil

	case http.MethodPatch:
		var (
			responseRaw []byte
			statusCode  int
			err         error
		)

		id, exists := req.PathParameters["id"]
		if !exists {
			return events.APIGatewayProxyResponse{
				StatusCode: http.StatusBadRequest,
				Body:       "id is required",
			}, nil
		}

		if req.Path == fmt.Sprintf("/orders/%s", id) {
			responseRaw, statusCode, err = ref.updateByID(ctx, []byte(req.Body), id)
		} else if req.Path == fmt.Sprintf("/orders/status/%s", id) {
			responseRaw, statusCode, err = ref.updateStatus(ctx, []byte(req.Body), id)
		} else {
			return events.APIGatewayProxyResponse{
				StatusCode: http.StatusNotFound,
			}, nil
		}

		response := string(responseRaw)

		if err != nil {
			response = err.Error()
		}

		return events.APIGatewayProxyResponse{
			StatusCode: statusCode,
			Headers: map[string]string{
				"Content-Type": "application-json",
			},
			Body: response,
		}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusMethodNotAllowed,
		Headers: map[string]string{
			"Content-Type": "application-json",
		},
		Body: errors.New("HTTP method not allowed").Error(),
	}, nil
}
