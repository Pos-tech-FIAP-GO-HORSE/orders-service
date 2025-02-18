package product_service

import (
	"context"

	"github.com/Pos-tech-FIAP-GO-HORSE/orders-service/src/core/domain/entity"
	"github.com/Pos-tech-FIAP-GO-HORSE/orders-service/src/infra/repository"
)

type ProductService struct {
	productRepository repository.IProductRepository
}

func NewProductService(productRepository repository.IProductRepository) IProductService {
	return &ProductService{productRepository}
}

func (ref *ProductService) Create(ctx context.Context, product entity.Product) (*entity.Product, error) {
	return ref.productRepository.Create(ctx, product)
}

func (ref *ProductService) Find(ctx context.Context) ([]*entity.Product, error) {
	return ref.productRepository.Find(ctx)
}

func (ref *ProductService) FindByID(ctx context.Context, id string) (*entity.Product, error) {
	return ref.productRepository.FindByID(ctx, id)
}

func (ref *ProductService) FindByPublicID(ctx context.Context, publicID string) (*entity.Product, error) {
	return ref.productRepository.FindByPublicID(ctx, publicID)
}

func (ref *ProductService) UpdateByID(ctx context.Context, id string, product entity.Product) (*entity.Product, error) {
	return ref.productRepository.UpdateByID(ctx, id, product)
}

func (ref *ProductService) DeleteByID(ctx context.Context, id string) (*entity.Product, error) {
	return ref.productRepository.DeleteByID(ctx, id)
}
