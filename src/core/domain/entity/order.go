package entity

import (
	"time"

	values "github.com/Pos-tech-FIAP-GO-HORSE/orders-service/src/core/domain/value_object"
)

type Order struct {
	ID                       string                 `json:"id"`
	Items                    []Item                 `json:"items"`
	TotalPrice               float64                `json:"totalPrice"`
	Status                   values.OrderStatusType `json:"status"`
	EstimatedPreparationTime int64                  `json:"estimatedPreparationTime"`
	CreatedAt                time.Time              `json:"createdAt"`
	UpdatedAt                time.Time              `json:"updatedAt"`
}

type Item struct {
	ID              string
	Name            string
	ImageURL        string
	Price           float64
	PreparationTime int64
	Quantity        int64
	Comments        string
}
