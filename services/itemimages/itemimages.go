package itemimages

import (
	"context"

	"github.com/baggyapp/depto"
	"github.com/baggyapp/depto/database"
	"github.com/baggyapp/depto/database/postgres"
	"github.com/baggyapp/depto/services"
)

type Config struct {
	DB postgres.SQLExecutor
}

func NewService(config Config) services.ItemImages {
	return service{
		store: postgres.NewItemImageStore(config.DB),
	}
}

type service struct {
	store database.ItemImageStore
}

func (svc service) GetByItemInfoID(ctx context.Context, itemInfoID int64) ([]*depto.ItemImage, error) {
	// only retrieve items that has not deleted.
	return svc.store.Select(ctx, depto.ItemImageQuery{
		ItemID:  itemInfoID,
		Deleted: false,
	})
}

func (svc service) Create(ctx context.Context, item *depto.ItemImage) error {
	return svc.store.Create(ctx, item)
}

func (svc service) Update(ctx context.Context, item *depto.ItemImage) error {
	return svc.store.Update(ctx, item)
}

func (svc service) Delete(ctx context.Context, item *depto.ItemImage) error {
	return svc.store.Delete(ctx, item)
}
