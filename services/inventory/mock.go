package inventory

import (
	"context"
	"time"

	"github.com/baggyapp/depto"
)

type MockService struct{}

func (svc MockService) GetByID(ctx context.Context, id string) (*depto.Inventory, error) {
	// only retrieve items that has not deleted.
	return &depto.Inventory{
		ID:        "1",
		Name:      "Test inventory",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		DeletedAt: nil,
	}, nil
}

func (svc MockService) Create(ctx context.Context, s *depto.Inventory) error {
	return nil
}

func (svc MockService) Update(ctx context.Context, s *depto.Inventory) error {
	return nil
}

func (svc MockService) Delete(ctx context.Context, s *depto.Inventory) error {
	return nil
}
