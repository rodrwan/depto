package inventoryitem

import (
	"context"

	"github.com/baggyapp/depto"
	"github.com/baggyapp/depto/database"
)

type Service struct {
	Store database.InventoryItemStore
}

func (svc Service) GetByID(ctx context.Context, id string) (*depto.InventoryItem, error) {
	// only retrieve items that has not deleted.
	return svc.Store.Get(ctx, depto.InventoryItemQuery{
		ID:      id,
		Deleted: false,
	})
}

func (svc Service) GetByInventoryID(ctx context.Context, inventoryID string) ([]*depto.InventoryItem, error) {
	// only retrieve items that has not deleted.
	return svc.Store.Select(ctx, depto.InventoryItemQuery{
		InventoryID: inventoryID,
		Deleted:     false,
	})
}

func (svc Service) Create(ctx context.Context, s *depto.InventoryItem) error {
	return svc.Store.Create(ctx, s)
}

func (svc Service) Update(ctx context.Context, s *depto.InventoryItem) error {
	return svc.Store.Update(ctx, s)
}

func (svc Service) Delete(ctx context.Context, s *depto.InventoryItem) error {
	return svc.Store.Delete(ctx, s)
}
