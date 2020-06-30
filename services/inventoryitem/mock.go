package inventoryitem

import (
	"context"
	"time"

	"github.com/baggyapp/depto"
)

type MockService struct {
}

func (svc MockService) GetByID(ctx context.Context, id string) (*depto.InventoryItem, error) {
	return &depto.InventoryItem{
		ID:          "1",
		InventoryID: "1",
		ItemID:      "1",
		Count:       2,
		LastUse:     nil,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		DeletedAt:   nil,
	}, nil
}

func (svc MockService) GetByInventoryID(ctx context.Context, inventoryID string) ([]*depto.InventoryItem, error) {
	// only retrieve items that has not deleted.
	return nil, nil
}

func (svc MockService) Create(ctx context.Context, s *depto.InventoryItem) error {
	return nil
}

func (svc MockService) Update(ctx context.Context, s *depto.InventoryItem) error {
	return nil
}

func (svc MockService) Delete(ctx context.Context, s *depto.InventoryItem) error {
	return nil
}
