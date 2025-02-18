package product_service

import (
	"context"
	"testing"

	"github.com/Pos-tech-FIAP-GO-HORSE/orders-service/src/core/domain/entity"
	"github.com/Pos-tech-FIAP-GO-HORSE/orders-service/src/infra/mocks"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNewProductService(t *testing.T) {
	expected := &ProductService{nil}
	actual := NewProductService(nil)
	assert.Equal(t, expected, actual)
}

func TestCreate(t *testing.T) {
	ctx := context.TODO()

	product := entity.Product{
		ID:              uuid.NewString(),
		Name:            "Batata frita",
		ImageUrl:        "batata_frita.png",
		Price:           4.99,
		PreparationTime: 3,
	}

	t.Run("should create a product", func(t *testing.T) {
		productRepositoryMocked := mocks.NewIProductRepository(t)
		productRepositoryMocked.
			On("Create", ctx, product).
			Return(&product, nil)

		service := ProductService{
			productRepository: productRepositoryMocked,
		}

		actual, err := service.Create(ctx, product)
		assert.Equal(t, &product, actual)
		assert.Nil(t, err)
	})
}

func TestFind(t *testing.T) {
	ctx := context.TODO()

	products := []*entity.Product{
		{
			ID:              uuid.NewString(),
			Name:            "Batata frita",
			ImageUrl:        "batata_frita.png",
			Price:           4.99,
			PreparationTime: 3,
		},
	}

	t.Run("should find products", func(t *testing.T) {
		productRepositoryMocked := mocks.NewIProductRepository(t)
		productRepositoryMocked.
			On("Find", ctx).
			Return(products, nil)

		service := ProductService{
			productRepository: productRepositoryMocked,
		}

		actual, err := service.Find(ctx)
		assert.Equal(t, products, actual)
		assert.Nil(t, err)
	})
}

func TestFindByID(t *testing.T) {
	ctx := context.TODO()

	product := entity.Product{
		ID:              uuid.NewString(),
		Name:            "Batata frita",
		ImageUrl:        "batata_frita.png",
		Price:           4.99,
		PreparationTime: 3,
	}

	t.Run("should find product", func(t *testing.T) {
		productRepositoryMocked := mocks.NewIProductRepository(t)
		productRepositoryMocked.
			On("FindByID", ctx, "id").
			Return(&product, nil)

		service := ProductService{
			productRepository: productRepositoryMocked,
		}

		actual, err := service.FindByID(ctx, "id")
		assert.Equal(t, &product, actual)
		assert.Nil(t, err)
	})
}

func TestFindByPublicID(t *testing.T) {
	ctx := context.TODO()

	product := entity.Product{
		ID:              uuid.NewString(),
		PublicID:        uuid.NewString(),
		Name:            "Batata frita",
		ImageUrl:        "batata_frita.png",
		Price:           4.99,
		PreparationTime: 3,
	}

	t.Run("should find product", func(t *testing.T) {
		productRepositoryMocked := mocks.NewIProductRepository(t)
		productRepositoryMocked.
			On("FindByPublicID", ctx, product.PublicID).
			Return(&product, nil)

		service := ProductService{
			productRepository: productRepositoryMocked,
		}

		actual, err := service.FindByPublicID(ctx, product.PublicID)
		assert.Equal(t, &product, actual)
		assert.Nil(t, err)
	})
}

func TestUpdateByID(t *testing.T) {
	ctx := context.TODO()

	product := entity.Product{
		ID:              uuid.NewString(),
		Name:            "Batata frita",
		ImageUrl:        "batata_frita.png",
		Price:           4.99,
		PreparationTime: 3,
	}

	t.Run("should update a product", func(t *testing.T) {
		productRepositoryMocked := mocks.NewIProductRepository(t)
		productRepositoryMocked.
			On("UpdateByID", ctx, "id", product).
			Return(&product, nil)

		service := ProductService{
			productRepository: productRepositoryMocked,
		}

		actual, err := service.UpdateByID(ctx, "id", product)
		assert.Equal(t, &product, actual)
		assert.Nil(t, err)
	})
}

func TestDeleteByID(t *testing.T) {
	ctx := context.TODO()

	product := entity.Product{
		ID:              uuid.NewString(),
		Name:            "Batata frita",
		ImageUrl:        "batata_frita.png",
		Price:           4.99,
		PreparationTime: 3,
	}

	t.Run("should delete a product", func(t *testing.T) {
		productRepositoryMocked := mocks.NewIProductRepository(t)
		productRepositoryMocked.
			On("DeleteByID", ctx, "id").
			Return(&product, nil)

		service := ProductService{
			productRepository: productRepositoryMocked,
		}

		actual, err := service.DeleteByID(ctx, "id")
		assert.Equal(t, &product, actual)
		assert.Nil(t, err)
	})
}
