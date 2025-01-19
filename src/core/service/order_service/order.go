package order_service

import (
	"context"

	"github.com/Pos-tech-FIAP-GO-HORSE/orders-service/src/core/domain/entity"
	"github.com/Pos-tech-FIAP-GO-HORSE/orders-service/src/infra/repository"
)

type OrderService struct {
	orderRepository repository.IOrderRepository
}

func NewOrderService(orderRepository repository.IOrderRepository) IOrderService {
	return &OrderService{orderRepository}
}

func (ref *OrderService) Create(ctx context.Context, order entity.Order) (*entity.Order, error) {
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
