package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/Pos-tech-FIAP-GO-HORSE/orders-service/src/api/contract"
)

func (ref *Handler) create(ctx context.Context, body []byte) ([]byte, int, error) {
	var orderRequest contract.CreateOrderRequest

	if err := json.Unmarshal(body, &orderRequest); err != nil {
		return nil, http.StatusBadRequest, err
	}

	order := orderRequest.ToDomain()

	createdOrder, err := ref.orderService.Create(ctx, order)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	response := contract.CreateOrderResponseFromDomain(createdOrder)
	rawResponse, _ := json.Marshal(response)

	return rawResponse, http.StatusCreated, nil
}

func (ref *Handler) findByID(ctx context.Context, id string) ([]byte, int, error) {
	order, err := ref.orderService.FindByID(ctx, id)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	response := contract.OrderFromDomain(*order)
	rawResponse, _ := json.Marshal(response)

	return rawResponse, http.StatusOK, nil
}

func (ref *Handler) find(ctx context.Context) ([]byte, int, error) {
	orders, err := ref.orderService.Find(ctx)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	response := contract.OrdersFromDomain(orders)
	rawResponse, _ := json.Marshal(response)

	return rawResponse, http.StatusOK, nil
}

func (ref *Handler) updateByID(ctx context.Context, body []byte, id string) ([]byte, int, error) {
	var orderRequest contract.UpdateOrderRequest

	if err := json.Unmarshal([]byte(body), &orderRequest); err != nil {
		return nil, http.StatusBadRequest, err
	}

	order := orderRequest.ToDomain()

	updatedOrder, err := ref.orderService.UpdateByID(ctx, id, order)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	response := contract.UpdateOrderResponseFromDomain(updatedOrder)
	rawResponse, _ := json.Marshal(response)

	return rawResponse, http.StatusOK, nil
}

func (ref *Handler) updateStatus(ctx context.Context, body []byte, id string) ([]byte, int, error) {
	var orderRequest contract.UpdateOrderStatusRequest

	if err := json.Unmarshal(body, &orderRequest); err != nil {
		return nil, http.StatusBadRequest, err
	}

	updatedOrder, err := ref.orderService.UpdateStatusByID(ctx, id, orderRequest.Status)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	response := contract.UpdateOrderResponseFromDomain(updatedOrder)
	rawResponse, _ := json.Marshal(response)

	return rawResponse, http.StatusOK, nil
}
