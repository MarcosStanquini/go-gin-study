package model

import "time"

type Product struct {
	ID        int       `json:"id_product"`
	Name      string    `json:"name"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"created_at"`
}