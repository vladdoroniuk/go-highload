package domain

import "time"

type Product struct {
	ID          int        `json:"id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Category    string     `json:"category"`
	Price       float64    `json:"price"`
	SKU         string     `json:"sku"`
	UpdatedAt   *time.Time `json:"updated_at"`
	CreatedAt   time.Time  `json:"created_at"`
}
