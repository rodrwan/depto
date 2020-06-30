package iteminfo

import (
	"context"
	"database/sql"

	"github.com/baggyapp/depto"
	"github.com/baggyapp/depto/database"
	"github.com/baggyapp/depto/database/postgres"
	"github.com/baggyapp/depto/services"
)

type Config struct {
	DB postgres.SQLExecutor
}

func NewService(config Config) services.ItemInfo {
	return service{
		store: postgres.NewItemInfoStore(config.DB),
	}
}

type service struct {
	store database.ItemInfoStore
}

func (svc service) GetByItemID(ctx context.Context, itemID string) (*depto.ItemInfo, error) {
	// only retrieve items that has not deleted.
	item, err := svc.store.Get(ctx, depto.ItemInfoQuery{
		ItemID:  itemID,
		Deleted: false,
	})
	if err == sql.ErrNoRows {
		return nil, nil
	}

	return item, nil
}

func (svc service) Create(ctx context.Context, item *depto.ItemInfo) error {
	return svc.store.Create(ctx, item)
}

func (svc service) Update(ctx context.Context, item *depto.ItemInfo) error {
	return svc.store.Update(ctx, item)
}

func (svc service) Delete(ctx context.Context, item *depto.ItemInfo) error {
	return svc.store.Delete(ctx, item)
}
