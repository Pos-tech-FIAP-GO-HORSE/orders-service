package order_service

import (
	"context"
	"encoding/json"

	"github.com/Pos-tech-FIAP-GO-HORSE/orders-service/src/core/domain/entity"
	values "github.com/Pos-tech-FIAP-GO-HORSE/orders-service/src/core/domain/value_object"
	"github.com/Pos-tech-FIAP-GO-HORSE/orders-service/src/infra/message_broker"
	"github.com/Pos-tech-FIAP-GO-HORSE/orders-service/src/infra/repository"
	"go.uber.org/zap"
)

type OrderService struct {
	orderRepository repository.IOrderRepository
	messageBroker   message_broker.IMessageBroker
	topics          map[string]string
}

func NewOrderService(orderRepository repository.IOrderRepository, messageBroker message_broker.IMessageBroker, topics map[string]string) IOrderService {
	return &OrderService{
		orderRepository,
		messageBroker,
		topics,
	}
}

func (ref *OrderService) Create(ctx context.Context, order entity.Order) (*entity.Order, error) {
	order.TotalPrice = ref.getTotalPrice(order.Items)
	order.EstimatedPreparationTime = ref.getEstimatedPreparationTime(order.Items)
	order.Status = values.TypeAwaitingPayment

	createdOrder, err := ref.orderRepository.Create(ctx, order)
	if err != nil {
		return nil, err
	}

	zap.L().Info("order created", zap.Any("order", createdOrder))

	messageRaw, _ := json.Marshal(createdOrder)
	if err = ref.messageBroker.Publish(ctx, ref.topics["order-created"], string(messageRaw)); err != nil {
		return nil, err
	}

	zap.L().Info("message sent to broker", zap.String("message", string(messageRaw)))

	return createdOrder, nil
}

func (ref *OrderService) Find(ctx context.Context) ([]*entity.Order, error) {
	return ref.orderRepository.Find(ctx)
}

func (ref *OrderService) FindByID(ctx context.Context, id string) (*entity.Order, error) {
	return ref.orderRepository.FindByID(ctx, id)
}

func (ref *OrderService) UpdateByID(ctx context.Context, id string, order entity.Order) (*entity.Order, error) {
	order.TotalPrice = ref.getTotalPrice(order.Items)
	order.EstimatedPreparationTime = ref.getEstimatedPreparationTime(order.Items)

	updatedOrder, err := ref.orderRepository.UpdateByID(ctx, id, order)
	if err != nil {
		return nil, err
	}

	zap.L().Info("order updated", zap.Any("order", updatedOrder))

	messageRaw, _ := json.Marshal(updatedOrder)
	if err = ref.messageBroker.Publish(ctx, ref.topics["order-updated"], string(messageRaw)); err != nil {
		return nil, err
	}

	zap.L().Info("message sent to broker", zap.String("message", string(messageRaw)))

	return updatedOrder, nil
}

func (ref *OrderService) UpdateStatusByID(ctx context.Context, id, status string) (*entity.Order, error) {
	updatedOrder, err := ref.orderRepository.UpdateStatusByID(ctx, id, status)
	if err != nil {
		return nil, err
	}

	zap.L().Info("order updated", zap.Any("order", updatedOrder))

	messageRaw, _ := json.Marshal(updatedOrder)
	if err = ref.messageBroker.Publish(ctx, ref.topics["order-updated"], string(messageRaw)); err != nil {
		return nil, err
	}

	zap.L().Info("message sent to broker", zap.String("message", string(messageRaw)))

	return updatedOrder, nil
}

func (ref *OrderService) getTotalPrice(items []entity.Item) float64 {
	var totalPrice float64

	for _, item := range items {
		totalPrice = totalPrice + (item.Price * float64(item.Quantity))
	}

	return totalPrice
}

func (ref *OrderService) getEstimatedPreparationTime(items []entity.Item) int64 {
	var estimatedPreparationTime int64

	for _, item := range items {
		estimatedPreparationTime = estimatedPreparationTime + (item.PreparationTime * item.Quantity)
	}

	return estimatedPreparationTime
}
