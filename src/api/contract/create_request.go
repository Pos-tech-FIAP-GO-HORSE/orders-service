package contract

import (
	"time"

	"github.com/Pos-tech-FIAP-GO-HORSE/orders-service/src/core/domain/entity"
)

type CreateOrderRequest struct {
	Items []Item `json:"items"`
}

type CreateOrderResponse struct {
	ID                       string    `json:"id"`
	PublicID                 string    `json:"public_id"`
	Items                    []Item    `json:"items"`
	TotalPrice               float64   `json:"total_price"`
	Status                   string    `json:"status"`
	EstimatedPreparationTime int64     `json:"estimated_preparation_time"`
	CreatedAt                time.Time `json:"createdAt"`
	UpdatedAt                time.Time `json:"updatedAt"`
}

func (ref CreateOrderRequest) ToDomain() entity.Order {
	items := make([]entity.Item, len(ref.Items))

	for i, item := range ref.Items {
		items[i] = item.ToDomain()
	}

	return entity.Order{
		Items: items,
	}
}

func CreateOrderResponseFromDomain(order *entity.Order) CreateOrderResponse {
	items := make([]Item, len(order.Items))

	for i, item := range order.Items {
		items[i] = ItemFromDomain(item)
	}

	return CreateOrderResponse{
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
