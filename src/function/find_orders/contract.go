package main

import (
	"time"

	"github.com/Pos-tech-FIAP-GO-HORSE/orders-service/src/core/domain/entity"
)

type Order struct {
	ID                       string    `json:"id"`
	Items                    []Item    `json:"items"`
	TotalPrice               float64   `json:"totalPrice"`
	Status                   string    `json:"status"`
	EstimatedPreparationTime int64     `json:"estimatedPreparationTime"`
	CreatedAt                time.Time `json:"createdAt"`
	UpdatedAt                time.Time `json:"updatedAt"`
}

type Orders []Order

type Item struct {
	ID              string  `json:"id"`
	Name            string  `json:"name"`
	ImageURL        string  `json:"imageUrl"`
	Price           float64 `json:"price"`
	Quantity        int64   `json:"quantity"`
	PreparationTime int64   `json:"preparationTime"`
	Comments        string  `json:"comments"`
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

func OrdersFromDomain(orders []*entity.Order) Orders {
	ordersResponse := make(Orders, len(orders))

	for i, order := range orders {
		items := make([]Item, len(order.Items))

		for j, item := range order.Items {
			items[j] = ItemFromDomain(item)
		}

		ordersResponse[i] = OrderFromDomain(*order)
		ordersResponse[i].Items = items
	}

	return ordersResponse
}

func OrderFromDomain(order entity.Order) Order {
	return Order{
		ID:                       order.ID,
		TotalPrice:               order.TotalPrice,
		Status:                   string(order.Status),
		EstimatedPreparationTime: order.EstimatedPreparationTime,
		CreatedAt:                order.CreatedAt,
		UpdatedAt:                order.UpdatedAt,
	}
}

func ItemFromDomain(item entity.Item) Item {
	return Item{
		ID:       item.ID,
		Price:    item.Price,
		Quantity: item.Quantity,
		Comments: item.Comments,
	}
}
