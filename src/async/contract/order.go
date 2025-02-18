package contract

import (
	"time"

	"github.com/Pos-tech-FIAP-GO-HORSE/orders-service/src/core/domain/entity"
)

type OrderResponse struct {
	ID                       string    `json:"id"`
	PublicID                 string    `json:"public_id"`
	Items                    []Item    `json:"items"`
	TotalPrice               float64   `json:"total_price"`
	Status                   string    `json:"status"`
	EstimatedPreparationTime int64     `json:"estimated_preparation_time"`
	CreatedAt                time.Time `json:"createdAt"`
	UpdatedAt                time.Time `json:"updatedAt"`
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

func OrderResponseFromDomain(order *entity.Order) OrderResponse {
	items := make([]Item, len(order.Items))

	for i, item := range order.Items {
		items[i] = ItemFromDomain(item)
	}

	return OrderResponse{
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

func ItemFromDomain(item entity.Item) Item {
	return Item{
		ID:       item.ID,
		PublicID: item.PublicID,
		Price:    item.Price,
		Quantity: item.Quantity,
		Comments: item.Comments,
	}
}
