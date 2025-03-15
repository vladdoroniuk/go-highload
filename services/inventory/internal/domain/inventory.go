package domain

import "time"

type Inventory struct {
	SKU               string    `json:"sku"`
	WarehouseID       string    `json:"warehouse_id"`
	QuantityAvailable int       `json:"quantity_available"`
	QuantityReserved  int       `json:"quantity_reserved"`
	LastUpdated       time.Time `json:"last_updated"`
}
