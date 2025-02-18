package handler

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"testing"

	"github.com/Pos-tech-FIAP-GO-HORSE/orders-service/src/api/contract"
	"github.com/Pos-tech-FIAP-GO-HORSE/orders-service/src/core/domain/entity"
	values "github.com/Pos-tech-FIAP-GO-HORSE/orders-service/src/core/domain/value_object"
	"github.com/Pos-tech-FIAP-GO-HORSE/orders-service/src/core/mocks"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	dummyError := errors.New("dummy error")

	orderRequest := contract.CreateOrderRequest{
		Items: []contract.Item{
			{
				ID:              uuid.NewString(),
				PublicID:        uuid.NewString(),
				Name:            "Batata Frita",
				ImageURL:        "batata-frita.png",
				Price:           5,
				PreparationTime: 2,
				Quantity:        1,
				Comments:        "",
			},
		},
	}

	rawOrder, _ := json.Marshal(orderRequest)

	createdOrder := &entity.Order{
		ID:       uuid.NewString(),
		PublicID: uuid.NewString(),
		Items: []entity.Item{
			{
				ID:              orderRequest.Items[0].ID,
				PublicID:        orderRequest.Items[0].PublicID,
				Name:            orderRequest.Items[0].Name,
				ImageURL:        orderRequest.Items[0].ImageURL,
				Price:           orderRequest.Items[0].Price,
				PreparationTime: orderRequest.Items[0].PreparationTime,
				Quantity:        orderRequest.Items[0].Quantity,
				Comments:        orderRequest.Items[0].Comments,
			},
		},
		TotalPrice:               5,
		Status:                   values.TypeAwaitingPayment,
		EstimatedPreparationTime: 2,
	}

	t.Run("should return an error when bad request", func(t *testing.T) {
		handler := Handler{}

		response, statusCode, err := handler.create(context.TODO(), []byte(`{"items": "test"}`))
		assert.Nil(t, response)
		assert.Equal(t, http.StatusBadRequest, statusCode)
		assert.Equal(t, "json: cannot unmarshal string into Go struct field CreateOrderRequest.items of type []contract.Item", err.Error())
	})

	t.Run("should return an error when create order", func(t *testing.T) {
		orderServiceMocked := mocks.NewIOrderService(t)
		orderServiceMocked.
			On("Create", context.TODO(), orderRequest.ToDomain()).
			Return(nil, dummyError)

		handler := Handler{orderService: orderServiceMocked}

		response, statusCode, err := handler.create(context.TODO(), rawOrder)
		assert.Nil(t, response)
		assert.Equal(t, http.StatusInternalServerError, statusCode)
		assert.Equal(t, "dummy error", err.Error())
	})

	t.Run("should create order successfully", func(t *testing.T) {
		orderServiceMocked := mocks.NewIOrderService(t)
		orderServiceMocked.
			On("Create", context.TODO(), orderRequest.ToDomain()).
			Return(createdOrder, nil)

		handler := Handler{orderService: orderServiceMocked}

		resp := contract.CreateOrderResponseFromDomain(createdOrder)
		rawResponse, _ := json.Marshal(resp)

		response, statusCode, err := handler.create(context.TODO(), rawOrder)
		assert.Equal(t, rawResponse, response)
		assert.Equal(t, http.StatusCreated, statusCode)
		assert.Nil(t, err)
	})
}

func TestFindByID(t *testing.T) {
	dummyError := errors.New("dummy error")

	order := &entity.Order{
		ID:       uuid.NewString(),
		PublicID: uuid.NewString(),
		Items: []entity.Item{
			{
				ID:              uuid.NewString(),
				PublicID:        uuid.NewString(),
				Name:            "name",
				ImageURL:        "image_url",
				Price:           5,
				PreparationTime: 1,
				Quantity:        1,
				Comments:        "",
			},
		},
		TotalPrice:               5,
		Status:                   values.TypeAwaitingPayment,
		EstimatedPreparationTime: 2,
	}

	t.Run("should return an error when find by id", func(t *testing.T) {
		orderServiceMocked := mocks.NewIOrderService(t)
		orderServiceMocked.
			On("FindByID", context.TODO(), "id").
			Return(nil, dummyError)

		handler := Handler{orderService: orderServiceMocked}

		response, statusCode, err := handler.findByID(context.TODO(), "id")
		assert.Nil(t, response)
		assert.Equal(t, http.StatusInternalServerError, statusCode)
		assert.Equal(t, dummyError, err)
	})

	t.Run("should find by id", func(t *testing.T) {
		orderServiceMocked := mocks.NewIOrderService(t)
		orderServiceMocked.
			On("FindByID", context.TODO(), "id").
			Return(order, nil)

		handler := Handler{orderService: orderServiceMocked}

		resp := contract.OrderFromDomain(*order)
		rawResponse, _ := json.Marshal(resp)

		response, statusCode, err := handler.findByID(context.TODO(), "id")
		assert.Equal(t, rawResponse, response)
		assert.Equal(t, http.StatusOK, statusCode)
		assert.Nil(t, err)
	})
}

func TestFind(t *testing.T) {
	dummyError := errors.New("dummy error")

	orders := []*entity.Order{
		{
			ID:       uuid.NewString(),
			PublicID: uuid.NewString(),
			Items: []entity.Item{
				{
					ID:              uuid.NewString(),
					PublicID:        uuid.NewString(),
					Name:            "name",
					ImageURL:        "image_url",
					Price:           5,
					PreparationTime: 1,
					Quantity:        1,
					Comments:        "",
				},
			},
			TotalPrice:               5,
			Status:                   values.TypeAwaitingPayment,
			EstimatedPreparationTime: 2,
		},
	}

	t.Run("should return an error when find", func(t *testing.T) {
		orderServiceMocked := mocks.NewIOrderService(t)
		orderServiceMocked.
			On("Find", context.TODO()).
			Return(nil, dummyError)

		handler := Handler{orderService: orderServiceMocked}

		response, statusCode, err := handler.find(context.TODO())
		assert.Nil(t, response)
		assert.Equal(t, http.StatusInternalServerError, statusCode)
		assert.Equal(t, dummyError, err)
	})

	t.Run("should find by id", func(t *testing.T) {
		orderServiceMocked := mocks.NewIOrderService(t)
		orderServiceMocked.
			On("Find", context.TODO()).
			Return(orders, nil)

		handler := Handler{orderService: orderServiceMocked}

		resp := contract.OrdersFromDomain(orders)
		rawResponse, _ := json.Marshal(resp)

		response, statusCode, err := handler.find(context.TODO())
		assert.Equal(t, rawResponse, response)
		assert.Equal(t, http.StatusOK, statusCode)
		assert.Nil(t, err)
	})
}

func TestUpdateByID(t *testing.T) {
	dummyError := errors.New("dummy error")

	orderRequest := contract.UpdateOrderRequest{
		Items: []contract.Item{
			{
				ID:              uuid.NewString(),
				PublicID:        uuid.NewString(),
				Name:            "Batata Frita",
				ImageURL:        "batata-frita.png",
				Price:           5,
				PreparationTime: 2,
				Quantity:        1,
				Comments:        "",
			},
		},
	}

	rawOrder, _ := json.Marshal(orderRequest)

	order := &entity.Order{
		ID:       uuid.NewString(),
		PublicID: uuid.NewString(),
		Items: []entity.Item{
			{
				ID:              uuid.NewString(),
				PublicID:        uuid.NewString(),
				Name:            "name",
				ImageURL:        "image_url",
				Price:           5,
				PreparationTime: 1,
				Quantity:        1,
				Comments:        "",
			},
		},
		TotalPrice:               5,
		Status:                   values.TypeAwaitingPayment,
		EstimatedPreparationTime: 2,
	}

	t.Run("should return an error when unmarshal", func(t *testing.T) {
		handler := Handler{}

		response, statusCode, err := handler.updateByID(context.TODO(), []byte(`{"items": "test"}`), "id")
		assert.Nil(t, response)
		assert.Equal(t, http.StatusBadRequest, statusCode)
		assert.Equal(t, "json: cannot unmarshal string into Go struct field UpdateOrderRequest.items of type []contract.Item", err.Error())
	})

	t.Run("should return an error when update by id", func(t *testing.T) {
		orderServiceMocked := mocks.NewIOrderService(t)
		orderServiceMocked.
			On("UpdateByID", context.TODO(), "id", orderRequest.ToDomain()).
			Return(nil, dummyError)

		handler := Handler{orderService: orderServiceMocked}

		response, statusCode, err := handler.updateByID(context.TODO(), rawOrder, "id")
		assert.Nil(t, response)
		assert.Equal(t, http.StatusInternalServerError, statusCode)
		assert.Equal(t, dummyError, err)
	})

	t.Run("should update by id successfully", func(t *testing.T) {
		orderServiceMocked := mocks.NewIOrderService(t)
		orderServiceMocked.
			On("UpdateByID", context.TODO(), "id", orderRequest.ToDomain()).
			Return(order, nil)

		handler := Handler{orderService: orderServiceMocked}

		resp := contract.UpdateOrderResponseFromDomain(order)
		rawResponse, _ := json.Marshal(resp)

		response, statusCode, err := handler.updateByID(context.TODO(), rawOrder, "id")
		assert.Equal(t, rawResponse, response)
		assert.Equal(t, http.StatusOK, statusCode)
		assert.Nil(t, err)
	})
}

func TestUpdateStatusByID(t *testing.T) {
	dummyError := errors.New("dummy error")

	orderRequest := contract.UpdateOrderStatusRequest{
		Status: string(values.TypeConfirmed),
	}

	rawOrder, _ := json.Marshal(orderRequest)

	order := &entity.Order{
		ID:       uuid.NewString(),
		PublicID: uuid.NewString(),
		Items: []entity.Item{
			{
				ID:              uuid.NewString(),
				PublicID:        uuid.NewString(),
				Name:            "name",
				ImageURL:        "image_url",
				Price:           5,
				PreparationTime: 1,
				Quantity:        1,
				Comments:        "",
			},
		},
		TotalPrice:               5,
		Status:                   values.TypeAwaitingPayment,
		EstimatedPreparationTime: 2,
	}

	t.Run("should return an error when unmarshal", func(t *testing.T) {
		handler := Handler{}

		response, statusCode, err := handler.updateStatus(context.TODO(), []byte(`{"status": 1}`), "id")
		assert.Nil(t, response)
		assert.Equal(t, http.StatusBadRequest, statusCode)
		assert.Equal(t, "json: cannot unmarshal number into Go struct field UpdateOrderStatusRequest.status of type string", err.Error())
	})

	t.Run("should return an error when update status by id", func(t *testing.T) {
		orderServiceMocked := mocks.NewIOrderService(t)
		orderServiceMocked.
			On("UpdateStatusByID", context.TODO(), "id", orderRequest.Status).
			Return(nil, dummyError)

		handler := Handler{orderService: orderServiceMocked}

		response, statusCode, err := handler.updateStatus(context.TODO(), rawOrder, "id")
		assert.Nil(t, response)
		assert.Equal(t, http.StatusInternalServerError, statusCode)
		assert.Equal(t, dummyError, err)
	})

	t.Run("should update status by id successfully", func(t *testing.T) {
		orderServiceMocked := mocks.NewIOrderService(t)
		orderServiceMocked.
			On("UpdateStatusByID", context.TODO(), "id", orderRequest.Status).
			Return(order, nil)

		handler := Handler{orderService: orderServiceMocked}

		resp := contract.UpdateOrderResponseFromDomain(order)
		rawResponse, _ := json.Marshal(resp)

		response, statusCode, err := handler.updateStatus(context.TODO(), rawOrder, "id")
		assert.Equal(t, rawResponse, response)
		assert.Equal(t, http.StatusOK, statusCode)
		assert.Nil(t, err)
	})
}
