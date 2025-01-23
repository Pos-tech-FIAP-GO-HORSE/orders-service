package repository

import (
	"context"

	"github.com/Pos-tech-FIAP-GO-HORSE/orders-service/src/core/domain/entity"
)

type IOrderRepository interface {
	Create(ctx context.Context, order entity.Order) (*entity.Order, error)
	Find(ctx context.Context) ([]*entity.Order, error)
	FindByID(ctx context.Context, id string) (*entity.Order, error)
	UpdateByID(ctx context.Context, id string, order entity.Order) (*entity.Order, error)
	UpdateStatusByID(ctx context.Context, id, status string) (*entity.Order, error)
}

type IProductRepository interface {
	Create(ctx context.Context, product entity.Product) (*entity.Product, error)
	Find(ctx context.Context) ([]*entity.Product, error)
	FindByID(ctx context.Context, id string) (*entity.Product, error)
	UpdateByID(ctx context.Context, id string, product entity.Product) (*entity.Product, error)
}
