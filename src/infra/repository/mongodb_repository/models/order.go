package models

import (
	"time"

	"github.com/Pos-tech-FIAP-GO-HORSE/orders-service/src/core/domain/entity"
	values "github.com/Pos-tech-FIAP-GO-HORSE/orders-service/src/core/domain/value_object"
)

type Order struct {
	ID                       string    `bson:"_id,omitempty"`
	Items                    []Item    `bson:"items,omitempty"`
	TotalPrice               float64   `bson:"totalPrice,omitempty"`
	Status                   string    `bson:"status,omitempty"`
	EstimatedPreparationTime int64     `bson:"estimatedPreparationTime,omitempty"`
	CreatedAt                time.Time `bson:"createdAt"`
	UpdatedAt                time.Time `bson:"updatedAt"`
}

type Item struct {
	ID              string  `bson:"id"`
	Name            string  `bson:"name,omitempty"`
	ImageURL        string  `bson:"imageUrl,omitempty"`
	Price           float64 `bson:"price,omitempty"`
	Quantity        int64   `bson:"quantity"`
	PreparationTime int64   `bson:"preparationTime,omitempty"`
	Comments        string  `bson:"comments,omitempty"`
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
		ID:              item.ID,
		Name:            item.Name,
		ImageURL:        item.ImageURL,
		Price:           item.Price,
		PreparationTime: item.PreparationTime,
		Quantity:        item.Quantity,
		Comments:        item.Comments,
	}
}

func (ref Item) ToDomain() entity.Item {
	return entity.Item{
		ID:              ref.ID,
		Name:            ref.Name,
		ImageURL:        ref.ImageURL,
		Price:           ref.Price,
		PreparationTime: ref.PreparationTime,
		Quantity:        ref.Quantity,
		Comments:        ref.Comments,
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
