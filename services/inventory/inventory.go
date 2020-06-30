package inventory

import (
	"context"

	"github.com/baggyapp/depto"
	"github.com/baggyapp/depto/database"
)

type Service struct {
	Store database.InventoryStore
}

func (svc Service) GetByID(ctx context.Context, id string) (*depto.Inventory, error) {
	// only retrieve items that has not deleted.
	return svc.Store.Get(ctx, depto.InventoryQuery{
		ID:      id,
		Deleted: false,
	})
}

func (svc Service) Create(ctx context.Context, s *depto.Inventory) error {
	return svc.Store.Create(ctx, s)
}

func (svc Service) Update(ctx context.Context, s *depto.Inventory) error {
	return svc.Store.Update(ctx, s)
}

func (svc Service) Delete(ctx context.Context, s *depto.Inventory) error {
	return svc.Store.Delete(ctx, s)
}
