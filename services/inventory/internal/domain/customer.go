package domain

import "time"

type Customer struct {
	ID        int64      `json:"id"`
	FirstName string     `json:"first_name"`
	LastName  string     `json:"last_name"`
	UpdatedAt *time.Time `json:"updated_at"`
	CreatedAt time.Time  `json:"created_at"`
}
