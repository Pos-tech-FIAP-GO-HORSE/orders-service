package product_service

import (
	"context"

	"github.com/Pos-tech-FIAP-GO-HORSE/orders-service/src/core/domain/entity"
)

type IProductService interface {
	Create(ctx context.Context, product entity.Product) (*entity.Product, error)
	Find(ctx context.Context) ([]*entity.Product, error)
	FindByID(ctx context.Context, id string) (*entity.Product, error)
	UpdateByID(ctx context.Context, id string, product entity.Product) (*entity.Product, error)
}
