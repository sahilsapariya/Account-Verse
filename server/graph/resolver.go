package graph

import (
	"server/database"
)

// This file will not be regenerated automatically.
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	DB *database.Database
}

func NewResolver(db *database.Database) *Resolver {
	return &Resolver{
		DB: db,
	}
}
