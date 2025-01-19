package models

import "github.com/Pos-tech-FIAP-GO-HORSE/orders-service/src/core/domain/entity"

type Product struct {
	ID              string  `bson:"_id,omitempty"`
	Name            string  `bson:"name,omitempty"`
	ImageUrl        string  `bson:"imageUrl,omitempty"`
	Price           float64 `bson:"price,omitempty"`
	PreparationTime int64   `bson:"preparationTime,omitempty"`
}

func ProductFromDomain(product entity.Product) Product {
	return Product{
		ID:              product.ID,
		Name:            product.Name,
		ImageUrl:        product.ImageUrl,
		Price:           product.Price,
		PreparationTime: product.PreparationTime,
	}
}

func (ref Product) ToDomain() entity.Product {
	return entity.Product{
		ID:              ref.ID,
		Name:            ref.Name,
		ImageUrl:        ref.ImageUrl,
		Price:           ref.Price,
		PreparationTime: ref.PreparationTime,
	}
}
