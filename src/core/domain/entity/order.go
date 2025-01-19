package entity

import (
	"time"

	values "github.com/Pos-tech-FIAP-GO-HORSE/orders-service/src/core/domain/value_object"
)

type Order struct {
	ID                       string
	Items                    []Item
	TotalPrice               float64
	Status                   values.OrderStatusType
	EstimatedPreparationTime int64
	CreatedAt                time.Time
	UpdatedAt                time.Time
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
