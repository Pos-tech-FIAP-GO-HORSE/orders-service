package order_service

import (
	"context"

	"github.com/Pos-tech-FIAP-GO-HORSE/orders-service/src/core/domain/entity"
	values "github.com/Pos-tech-FIAP-GO-HORSE/orders-service/src/core/domain/value_object"
	"github.com/Pos-tech-FIAP-GO-HORSE/orders-service/src/infra/repository"
)

type OrderService struct {
	orderRepository repository.IOrderRepository
}

func NewOrderService(orderRepository repository.IOrderRepository) IOrderService {
	return &OrderService{orderRepository}
}

func (ref *OrderService) Create(ctx context.Context, order entity.Order) (*entity.Order, error) {
	order.TotalPrice = ref.getTotalPrice(order.Items)
	order.EstimatedPreparationTime = ref.getEstimatedPreparationTime(order.Items)
	order.Status = values.TypeAwaitingPayment

	return ref.orderRepository.Create(ctx, order)
}

func (ref *OrderService) Find(ctx context.Context) ([]*entity.Order, error) {
	return ref.orderRepository.Find(ctx)
}

func (ref *OrderService) FindByID(ctx context.Context, id string) (*entity.Order, error) {
	return ref.orderRepository.FindByID(ctx, id)
}

func (ref *OrderService) UpdateByID(ctx context.Context, id string, order entity.Order) (*entity.Order, error) {
	return ref.orderRepository.UpdateByID(ctx, id, order)
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
