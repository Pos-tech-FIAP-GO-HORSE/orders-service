package order_service

import (
	"context"
	"encoding/json"
	"errors"
	"testing"
	"time"

	"github.com/Pos-tech-FIAP-GO-HORSE/orders-service/src/core/domain/entity"
	values "github.com/Pos-tech-FIAP-GO-HORSE/orders-service/src/core/domain/value_object"
	"github.com/Pos-tech-FIAP-GO-HORSE/orders-service/src/infra/mocks"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNewOrderService(t *testing.T) {
	expected := &OrderService{
		nil, nil, map[string]string{},
	}
	actual := NewOrderService(nil, nil, map[string]string{})
	assert.Equal(t, expected, actual)
}

func TestCreate(t *testing.T) {
	ctx := context.TODO()
	unexpectedError := errors.New("unexpected error")

	order := entity.Order{
		ID:     uuid.NewString(),
		Status: values.TypeAwaitingPayment,
		Items: []entity.Item{
			{
				ID:              uuid.NewString(),
				Name:            "Batata frita",
				ImageURL:        "batata_frita.png",
				Price:           4.99,
				PreparationTime: 3,
				Quantity:        1,
				Comments:        "",
			},
		},
		TotalPrice:               4.99,
		EstimatedPreparationTime: 3,
		CreatedAt:                time.Now(),
		UpdatedAt:                time.Now(),
	}

	t.Run("should not create an order when error to persist on database", func(t *testing.T) {
		orderRepositoryMocked := mocks.NewIOrderRepository(t)
		orderRepositoryMocked.
			On("Create", ctx, order).
			Return(nil, unexpectedError)

		service := OrderService{
			orderRepository: orderRepositoryMocked,
		}

		actual, err := service.Create(ctx, order)
		assert.Equal(t, unexpectedError, err)
		assert.Nil(t, actual)
	})

	t.Run("should not create an order when error to send message", func(t *testing.T) {
		orderRepositoryMocked := mocks.NewIOrderRepository(t)
		orderRepositoryMocked.
			On("Create", ctx, order).
			Return(&order, nil)

		messageRaw, _ := json.Marshal(order)
		messageBrokerMocked := mocks.NewIMessageBroker(t)
		messageBrokerMocked.
			On("Publish", ctx, "order-created", string(messageRaw)).
			Return(unexpectedError)

		service := OrderService{
			orderRepository: orderRepositoryMocked,
			messageBroker:   messageBrokerMocked,
			topics: map[string]string{
				"order-created": "order-created",
			},
		}

		actual, err := service.Create(ctx, order)
		assert.Equal(t, unexpectedError, err)
		assert.Nil(t, actual)
	})

	t.Run("should create an order successfully", func(t *testing.T) {
		orderRepositoryMocked := mocks.NewIOrderRepository(t)
		orderRepositoryMocked.
			On("Create", ctx, order).
			Return(&order, nil)

		messageRaw, _ := json.Marshal(order)
		messageBrokerMocked := mocks.NewIMessageBroker(t)
		messageBrokerMocked.
			On("Publish", ctx, "order-created", string(messageRaw)).
			Return(nil)

		service := OrderService{
			orderRepository: orderRepositoryMocked,
			messageBroker:   messageBrokerMocked,
			topics: map[string]string{
				"order-created": "order-created",
			},
		}

		actual, err := service.Create(ctx, order)
		assert.Equal(t, &order, actual)
		assert.Nil(t, err)
	})
}

func TestFind(t *testing.T) {
	ctx := context.TODO()

	orders := []*entity.Order{
		{
			ID:     uuid.NewString(),
			Status: values.TypeAwaitingPayment,
			Items: []entity.Item{
				{
					ID:              uuid.NewString(),
					Name:            "Batata frita",
					ImageURL:        "batata_frita.png",
					Price:           4.99,
					PreparationTime: 3,
					Quantity:        1,
					Comments:        "",
				},
			},
			TotalPrice:               4.99,
			EstimatedPreparationTime: 3,
			CreatedAt:                time.Now(),
			UpdatedAt:                time.Now(),
		},
	}

	t.Run("should find orders successfully", func(t *testing.T) {
		orderRepositoryMocked := mocks.NewIOrderRepository(t)
		orderRepositoryMocked.On("Find", ctx).Return(orders, nil)

		service := OrderService{
			orderRepository: orderRepositoryMocked,
		}

		actual, err := service.Find(ctx)
		assert.Equal(t, orders, actual)
		assert.Nil(t, err)
	})
}

func TestFindByID(t *testing.T) {
	ctx := context.TODO()

	order := &entity.Order{
		ID:     uuid.NewString(),
		Status: values.TypeAwaitingPayment,
		Items: []entity.Item{
			{
				ID:              uuid.NewString(),
				Name:            "Batata frita",
				ImageURL:        "batata_frita.png",
				Price:           4.99,
				PreparationTime: 3,
				Quantity:        1,
				Comments:        "",
			},
		},
		TotalPrice:               4.99,
		EstimatedPreparationTime: 3,
		CreatedAt:                time.Now(),
		UpdatedAt:                time.Now(),
	}

	t.Run("should find orders successfully", func(t *testing.T) {
		orderRepositoryMocked := mocks.NewIOrderRepository(t)
		orderRepositoryMocked.On("FindByID", ctx, "id").Return(order, nil)

		service := OrderService{
			orderRepository: orderRepositoryMocked,
		}

		actual, err := service.FindByID(ctx, "id")
		assert.Equal(t, order, actual)
		assert.Nil(t, err)
	})
}

func TestUpdateByID(t *testing.T) {
	ctx := context.TODO()
	unexpectedError := errors.New("unexpected error")

	order := entity.Order{
		ID:     uuid.NewString(),
		Status: values.TypeAwaitingPayment,
		Items: []entity.Item{
			{
				ID:              uuid.NewString(),
				Name:            "Batata frita",
				ImageURL:        "batata_frita.png",
				Price:           4.99,
				PreparationTime: 3,
				Quantity:        1,
				Comments:        "",
			},
		},
		TotalPrice:               4.99,
		EstimatedPreparationTime: 3,
		CreatedAt:                time.Now(),
		UpdatedAt:                time.Now(),
	}

	t.Run("should not update an order when error to persist on database", func(t *testing.T) {
		orderRepositoryMocked := mocks.NewIOrderRepository(t)
		orderRepositoryMocked.
			On("UpdateByID", ctx, "id", order).
			Return(nil, unexpectedError)

		service := OrderService{
			orderRepository: orderRepositoryMocked,
		}

		actual, err := service.UpdateByID(ctx, "id", order)
		assert.Equal(t, unexpectedError, err)
		assert.Nil(t, actual)
	})

	t.Run("should not update an order when error to send message", func(t *testing.T) {
		orderRepositoryMocked := mocks.NewIOrderRepository(t)
		orderRepositoryMocked.
			On("UpdateByID", ctx, "id", order).
			Return(&order, nil)

		messageRaw, _ := json.Marshal(order)
		messageBrokerMocked := mocks.NewIMessageBroker(t)
		messageBrokerMocked.
			On("Publish", ctx, "order-updated", string(messageRaw)).
			Return(unexpectedError)

		service := OrderService{
			orderRepository: orderRepositoryMocked,
			messageBroker:   messageBrokerMocked,
			topics: map[string]string{
				"order-updated": "order-updated",
			},
		}

		actual, err := service.UpdateByID(ctx, "id", order)
		assert.Equal(t, unexpectedError, err)
		assert.Nil(t, actual)
	})

	t.Run("should update an order successfully", func(t *testing.T) {
		orderRepositoryMocked := mocks.NewIOrderRepository(t)
		orderRepositoryMocked.
			On("UpdateByID", ctx, "id", order).
			Return(&order, nil)

		messageRaw, _ := json.Marshal(order)
		messageBrokerMocked := mocks.NewIMessageBroker(t)
		messageBrokerMocked.
			On("Publish", ctx, "order-updated", string(messageRaw)).
			Return(nil)

		service := OrderService{
			orderRepository: orderRepositoryMocked,
			messageBroker:   messageBrokerMocked,
			topics: map[string]string{
				"order-updated": "order-updated",
			},
		}

		actual, err := service.UpdateByID(ctx, "id", order)
		assert.Equal(t, &order, actual)
		assert.Nil(t, err)
	})
}

func TestUpdateStatusByID(t *testing.T) {
	ctx := context.TODO()
	unexpectedError := errors.New("unexpected error")

	order := entity.Order{
		ID:     uuid.NewString(),
		Status: values.TypeAwaitingPayment,
		Items: []entity.Item{
			{
				ID:              uuid.NewString(),
				Name:            "Batata frita",
				ImageURL:        "batata_frita.png",
				Price:           4.99,
				PreparationTime: 3,
				Quantity:        1,
				Comments:        "",
			},
		},
		TotalPrice:               4.99,
		EstimatedPreparationTime: 3,
		CreatedAt:                time.Now(),
		UpdatedAt:                time.Now(),
	}

	t.Run("should not update an order status when error to persist on database", func(t *testing.T) {
		orderRepositoryMocked := mocks.NewIOrderRepository(t)
		orderRepositoryMocked.
			On("UpdateStatusByID", ctx, "id", string(values.TypeConfirmed)).
			Return(nil, unexpectedError)

		service := OrderService{
			orderRepository: orderRepositoryMocked,
		}

		actual, err := service.UpdateStatusByID(ctx, "id", string(values.TypeConfirmed))
		assert.Equal(t, unexpectedError, err)
		assert.Nil(t, actual)
	})

	t.Run("should not update an order status when error to send message", func(t *testing.T) {
		orderRepositoryMocked := mocks.NewIOrderRepository(t)
		orderRepositoryMocked.
			On("UpdateStatusByID", ctx, "id", string(values.TypeConfirmed)).
			Return(&order, nil)

		messageRaw, _ := json.Marshal(order)
		messageBrokerMocked := mocks.NewIMessageBroker(t)
		messageBrokerMocked.
			On("Publish", ctx, "order-updated", string(messageRaw)).
			Return(unexpectedError)

		service := OrderService{
			orderRepository: orderRepositoryMocked,
			messageBroker:   messageBrokerMocked,
			topics: map[string]string{
				"order-updated": "order-updated",
			},
		}

		actual, err := service.UpdateStatusByID(ctx, "id", string(values.TypeConfirmed))
		assert.Equal(t, unexpectedError, err)
		assert.Nil(t, actual)
	})

	t.Run("should update an order successfully", func(t *testing.T) {
		orderRepositoryMocked := mocks.NewIOrderRepository(t)
		orderRepositoryMocked.
			On("UpdateStatusByID", ctx, "id", string(values.TypeConfirmed)).
			Return(&order, nil)

		messageRaw, _ := json.Marshal(order)
		messageBrokerMocked := mocks.NewIMessageBroker(t)
		messageBrokerMocked.
			On("Publish", ctx, "order-updated", string(messageRaw)).
			Return(nil)

		service := OrderService{
			orderRepository: orderRepositoryMocked,
			messageBroker:   messageBrokerMocked,
			topics: map[string]string{
				"order-updated": "order-updated",
			},
		}

		actual, err := service.UpdateStatusByID(ctx, "id", string(values.TypeConfirmed))
		assert.Equal(t, &order, actual)
		assert.Nil(t, err)
	})
}
