package app

import (
	"github.com/go-highload/services/inventory/internal/config"
	"github.com/go-highload/services/inventory/internal/domain"
	"github.com/go-highload/services/inventory/pkg/cassandra"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

func Run() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	e := echo.New()

	db, err := cassandra.NewClient(cfg.Cassandra.User, cfg.Cassandra.Password, cfg.Cassandra.Host, cfg.Cassandra.Port, cfg.Cassandra.Keyspace)
	if err != nil {
		log.Fatalf("failed to connect to Cassandra")
	}

	e.GET("/", func(c echo.Context) error {
		q := "SELECT sku, warehouse_id, quantity_available, quantity_reserved, last_updated FROM inventory_by_sku"

		iter := db.Query(q).Iter()

		var inventory []domain.Inventory
		for {
			var i domain.Inventory
			if !iter.Scan(&i.SKU, &i.WarehouseID, &i.QuantityAvailable, &i.QuantityReserved, &i.LastUpdated) {
				break
			}
			inventory = append(inventory, i)
		}

		if err := iter.Close(); err != nil {
			return c.JSON(http.StatusInternalServerError, "Error reading inventory")
		}

		return c.JSON(http.StatusOK, struct {
			Inventory []domain.Inventory `json:"inventory"`
		}{
			Inventory: inventory,
		})
	})

	e.Logger.Fatal(e.Start(":" + cfg.App.Port))
}
