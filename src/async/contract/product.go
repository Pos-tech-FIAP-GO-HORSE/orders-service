package contract

import (
	"time"

	"github.com/Pos-tech-FIAP-GO-HORSE/orders-service/src/core/domain/entity"
)

type ProductEvent struct {
	ID              string    `json:"id"`
	PublicID        string    `json:"public_id"`
	Name            string    `json:"name"`
	Category        string    `json:"category"`
	Price           float64   `json:"price"`
	Description     string    `json:"description"`
	ImageUrl        string    `json:"image_url"`
	IsAvailable     bool      `json:"is_available"`
	PreparationTime int64     `json:"preparation_time"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

func (ref ProductEvent) ToDomain() entity.Product {
	return entity.Product{
		ID:              ref.ID,
		PublicID:        ref.PublicID,
		Name:            ref.Name,
		ImageUrl:        ref.ImageUrl,
		Price:           ref.Price,
		PreparationTime: ref.PreparationTime,
	}
}
