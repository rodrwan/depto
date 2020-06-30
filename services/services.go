package services

import (
	"context"

	"github.com/baggyapp/depto"
)

type Items interface {
	GetByID(ctx context.Context, id string) (*depto.Item, error)
	GetByName(ctx context.Context, name string) (*depto.Item, error)
	Select(ctx context.Context) ([]*depto.Item, error)
	Create(ctx context.Context, i *depto.Item) error
	Update(ctx context.Context, id string, i *depto.Item) (*depto.Item, error)
	Delete(ctx context.Context, i *depto.Item) error
}

type ItemImages interface {
	GetByItemInfoID(ctx context.Context, itemInfoID int64) ([]*depto.ItemImage, error)
	Create(ctx context.Context, i *depto.ItemImage) error
	Update(ctx context.Context, i *depto.ItemImage) error
	Delete(ctx context.Context, i *depto.ItemImage) error
}

type ItemInfo interface {
	GetByItemID(ctx context.Context, itemID string) (*depto.ItemInfo, error)
	Create(ctx context.Context, i *depto.ItemInfo) error
	Update(ctx context.Context, i *depto.ItemInfo) error
	Delete(ctx context.Context, i *depto.ItemInfo) error
}

type Inventory interface {
	GetByID(ctx context.Context, id string) (*depto.Inventory, error)
	Create(ctx context.Context, s *depto.Inventory) error
	Update(ctx context.Context, s *depto.Inventory) error
	Delete(ctx context.Context, s *depto.Inventory) error
}

type InventoryItem interface {
	GetByID(ctx context.Context, id string) (*depto.InventoryItem, error)
	GetByInventoryID(ctx context.Context, inventoryID string) ([]*depto.InventoryItem, error)
	Create(ctx context.Context, s *depto.InventoryItem) error
	Update(ctx context.Context, s *depto.InventoryItem) error
	Delete(ctx context.Context, s *depto.InventoryItem) error
}

type Repositions interface {
	GetByID(ctx context.Context, id string) (*depto.Repositions, error)
	GetByCollectionID(ctx context.Context, collectionID string) (*depto.Repositions, error)
	Create(ctx context.Context, r *depto.Repositions) error
	Update(ctx context.Context, r *depto.Repositions) error
	Delete(ctx context.Context, r *depto.Repositions) error
}
