package models

import (
	"time"

	"github.com/Pos-tech-FIAP-GO-HORSE/orders-service/src/core/domain/entity"
	values "github.com/Pos-tech-FIAP-GO-HORSE/orders-service/src/core/domain/value_object"
)

type Order struct {
	ID                       string    `bson:"_id,omitempty"`
	Items                    []Item    `bson:"items"`
	TotalPrice               float64   `bson:"totalPrice"`
	Status                   string    `bson:"status"`
	EstimatedPreparationTime int64     `bson:"estimatedPreparationTime"`
	CreatedAt                time.Time `bson:"createdAt"`
	UpdatedAt                time.Time `bson:"updatedAt"`
}

type Item struct {
	Product
	Comments string `bson:"comments"`
}

func OrderFromDomain(order entity.Order) Order {
	items := make([]Item, len(order.Items))

	for i, item := range order.Items {
		items[i] = ItemFromDomain(item)
	}

	return Order{
		ID:                       order.ID,
		Items:                    items,
		TotalPrice:               order.TotalPrice,
		Status:                   string(order.Status),
		EstimatedPreparationTime: order.EstimatedPreparationTime,
	}
}

func ItemFromDomain(item entity.Item) Item {
	return Item{
		Product: Product{
			ID:              item.ID,
			Name:            item.Name,
			ImageUrl:        item.ImageUrl,
			Price:           item.Price,
			PreparationTime: item.PreparationTime,
		},
		Comments: item.Comments,
	}
}

func (ref Item) ToDomain() entity.Item {
	return entity.Item{
		Product: entity.Product{
			ID:              ref.ID,
			Name:            ref.Name,
			ImageUrl:        ref.ImageUrl,
			Price:           ref.Price,
			PreparationTime: ref.PreparationTime,
		},
		Comments: ref.Comments,
	}
}

func (ref Order) ToDomain() entity.Order {
	items := make([]entity.Item, len(ref.Items))

	for i, item := range ref.Items {
		items[i] = item.ToDomain()
	}

	return entity.Order{
		ID:                       ref.ID,
		Items:                    items,
		TotalPrice:               ref.TotalPrice,
		Status:                   values.OrderStatusType(ref.Status),
		EstimatedPreparationTime: ref.EstimatedPreparationTime,
		CreatedAt:                ref.CreatedAt,
		UpdatedAt:                ref.UpdatedAt,
	}
}
