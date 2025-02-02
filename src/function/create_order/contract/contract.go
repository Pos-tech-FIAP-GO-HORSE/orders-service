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
	Items                    []Item    `json:"items"`
	TotalPrice               float64   `json:"totalPrice"`
	Status                   string    `json:"status"`
	EstimatedPreparationTime int64     `json:"estimatedPreparationTime"`
	CreatedAt                time.Time `json:"createdAt"`
	UpdatedAt                time.Time `json:"updatedAt"`
}

type Item struct {
	ID              string  `json:"id"`
	Name            string  `json:"name"`
	ImageURL        string  `json:"imageUrl"`
	Price           float64 `json:"price"`
	Quantity        int64   `json:"quantity"`
	PreparationTime int64   `json:"preparationTime"`
	Comments        string  `json:"comments"`
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

func CreateOrderResponseFromDomain(order *entity.Order) CreateOrderResponse {
	items := make([]Item, len(order.Items))

	for i, item := range order.Items {
		items[i] = ItemFromDomain(item)
	}

	return CreateOrderResponse{
		ID:                       order.ID,
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
		Price:    item.Price,
		Quantity: item.Quantity,
		Comments: item.Comments,
	}
}
