package app

import (
	"context"
	"github.com/go-highload/services/products/internal/config"
	"github.com/go-highload/services/products/internal/domain"
	"github.com/go-highload/services/products/pkg/elasticsearch"
	"github.com/go-highload/services/products/pkg/postgres"
	"github.com/labstack/echo/v4"
	"log"
	"log/slog"
	"net/http"
)

func Run() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	e := echo.New()

	db, err := postgres.NewClient(cfg.Postgres.User, cfg.Postgres.Password, cfg.Postgres.Host, cfg.Postgres.Port, cfg.Postgres.DB)
	if err != nil {
		log.Fatalf("failed to connect to Postgres")
	}

	es, err := elasticsearch.NewClient(cfg.Elasticsearch.User, cfg.Elasticsearch.Password, cfg.Elasticsearch.Host, cfg.Elasticsearch.Port)
	if err != nil {
		log.Fatalf("failed to connect to Elasticsearch")
	}

	e.GET("/", func(c echo.Context) error {
		q := "SELECT * FROM products"

		rows, err := db.Query(c.Request().Context(), q)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, "Error querying database")
		}
		defer rows.Close()

		var customers []domain.Product
		for rows.Next() {
			var cu domain.Product
			if err := rows.Scan(&cu.ID, &cu.FirstName, &cu.LastName, &cu.UpdatedAt, &cu.CreatedAt); err != nil {
				slog.Log(context.Background(), slog.LevelError, "Error processing data", err)
				return c.JSON(http.StatusInternalServerError, "Error processing data")
			}
			customers = append(customers, cu)
		}

		if err := rows.Err(); err != nil {
			return c.JSON(http.StatusInternalServerError, "Error reading customers")
		}

		return c.JSON(http.StatusOK, struct {
			Products []domain.Product `json:"customers"`
		}{
			Products: customers,
		})
	})

	e.Logger.Fatal(e.Start(":" + cfg.App.Port))
}
