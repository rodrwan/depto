package database

import (
	"context"

	"github.com/baggyapp/depto"
	"github.com/jmoiron/sqlx"
)

// ItemStore define CRUD functions that need to be defined by any structure to operate over Item table.
type ItemStore interface {
	Get(ctx context.Context, isq depto.ItemQuery) (*depto.Item, error)
	Select(ctx context.Context, isq depto.ItemQuery) ([]*depto.Item, error)
	Create(ctx context.Context, i *depto.Item) error
	Update(ctx context.Context, i *depto.Item) error
	Delete(ctx context.Context, i *depto.Item) error
}

// ItemInfoStore define CRUD functions that need to be defined by any structure to operate over ItemInfo table.
type ItemInfoStore interface {
	Get(ctx context.Context, isq depto.ItemInfoQuery) (*depto.ItemInfo, error)
	Select(ctx context.Context, isq depto.ItemInfoQuery) ([]*depto.ItemInfo, error)
	Create(ctx context.Context, i *depto.ItemInfo) error
	Update(ctx context.Context, i *depto.ItemInfo) error
	Delete(ctx context.Context, i *depto.ItemInfo) error
}

// ItemImageStore define CRUD functions that need to be defined by any structure to operate over ItemInfo table.
type ItemImageStore interface {
	Get(ctx context.Context, isq depto.ItemImageQuery) (*depto.ItemImage, error)
	Select(ctx context.Context, isq depto.ItemImageQuery) ([]*depto.ItemImage, error)
	Create(ctx context.Context, i *depto.ItemImage) error
	Update(ctx context.Context, i *depto.ItemImage) error
	Delete(ctx context.Context, i *depto.ItemImage) error
}

// InventoryStore define CRUD functions that need to be defined by any structure to operate over Stock table.
type InventoryStore interface {
	Get(ctx context.Context, isq depto.InventoryQuery) (*depto.Inventory, error)
	Select(ctx context.Context, isq depto.InventoryQuery) ([]*depto.Inventory, error)
	Create(ctx context.Context, i *depto.Inventory) error
	Update(ctx context.Context, i *depto.Inventory) error
	Delete(ctx context.Context, i *depto.Inventory) error
}

// InventoryItemStore define CRUD functions that need to be defined by any structure to operate over Stock table.
type InventoryItemStore interface {
	Get(ctx context.Context, isq depto.InventoryItemQuery) (*depto.InventoryItem, error)
	Select(ctx context.Context, isq depto.InventoryItemQuery) ([]*depto.InventoryItem, error)
	Create(ctx context.Context, i *depto.InventoryItem) error
	Update(ctx context.Context, i *depto.InventoryItem) error
	Delete(ctx context.Context, i *depto.InventoryItem) error
}

// RepositionsStore define CRUD functions that need to be defined by any structure to operate over Repositions table.
type RepositionsStore interface {
	Get(ctx context.Context, isq depto.RepositionsQuery) (*depto.Repositions, error)
	Select(ctx context.Context, isq depto.RepositionsQuery) ([]*depto.Repositions, error)
	Create(ctx context.Context, i *depto.Repositions) error
	Update(ctx context.Context, i *depto.Repositions) error
	Delete(ctx context.Context, i *depto.Repositions) error
}

// NewPostgres create a new connection to postgres database
func NewPostgres(dsn string) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
