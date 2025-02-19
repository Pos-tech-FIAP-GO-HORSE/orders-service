package entity

import (
	"time"

	values "github.com/Pos-tech-FIAP-GO-HORSE/orders-service/src/core/domain/value_object"
)

type Order struct {
	ID                       string                 `json:"order_id"`
	PublicID                 string                 `json:"public_id"`
	Items                    []Item                 `json:"items"`
	TotalPrice               float64                `json:"amount"`
	Status                   values.OrderStatusType `json:"status"`
	EstimatedPreparationTime int64                  `json:"estimatedPreparationTime"`
	CreatedAt                time.Time              `json:"createdAt"`
	UpdatedAt                time.Time              `json:"updatedAt"`
}

type Item struct {
	ID              string  `json:"id"`
	PublicID        string  `json:"public_id"`
	Name            string  `json:"name"`
	ImageURL        string  `json:"imageURL"`
	Price           float64 `json:"price"`
	PreparationTime int64   `json:"preparationTime"`
	Quantity        int64   `json:"quantity"`
	Comments        string  `json:"comments"`
}
