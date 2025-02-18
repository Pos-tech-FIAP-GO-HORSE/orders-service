package contract

import (
	"time"

	"github.com/Pos-tech-FIAP-GO-HORSE/orders-service/src/core/domain/entity"
)

type Order struct {
	ID                       string    `json:"id"`
	PublicID                 string    `json:"public_id"`
	Items                    []Item    `json:"items"`
	TotalPrice               float64   `json:"total_price"`
	Status                   string    `json:"status"`
	EstimatedPreparationTime int64     `json:"estimated_preparation_time"`
	CreatedAt                time.Time `json:"created_at"`
	UpdatedAt                time.Time `json:"updated_at"`
}

type Orders []Order

func OrderFromDomain(order entity.Order) Order {
	items := make([]Item, len(order.Items))

	for i, item := range order.Items {
		items[i] = ItemFromDomain(item)
	}

	return Order{
		ID:                       order.ID,
		PublicID:                 order.PublicID,
		Items:                    items,
		TotalPrice:               order.TotalPrice,
		Status:                   string(order.Status),
		EstimatedPreparationTime: order.EstimatedPreparationTime,
		CreatedAt:                order.CreatedAt,
		UpdatedAt:                order.UpdatedAt,
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

type Item struct {
	ID              string  `json:"id"`
	PublicID        string  `json:"public_id"`
	Name            string  `json:"name"`
	ImageURL        string  `json:"image_url"`
	Price           float64 `json:"price"`
	Quantity        int64   `json:"quantity"`
	PreparationTime int64   `json:"preparation_time"`
	Comments        string  `json:"comments"`
}

func (ref Item) ToDomain() entity.Item {
	return entity.Item{
		ID:              ref.ID,
		PublicID:        ref.PublicID,
		Name:            ref.Name,
		ImageURL:        ref.ImageURL,
		Price:           ref.Price,
		PreparationTime: ref.PreparationTime,
		Quantity:        ref.Quantity,
		Comments:        ref.Comments,
	}
}

func ItemFromDomain(item entity.Item) Item {
	return Item{
		ID:       item.ID,
		PublicID: item.PublicID,
		Price:    item.Price,
		Quantity: item.Quantity,
		Comments: item.Comments,
	}
}
