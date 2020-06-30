package items

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

func NewService(config Config) services.Items {
	return service{
		store: postgres.NewItemsStore(config.DB),
	}
}

type service struct {
	store database.ItemStore
}

func (svc service) GetByID(ctx context.Context, id string) (*depto.Item, error) {
	// only retrieve items that has not deleted.
	return svc.store.Get(ctx, depto.ItemQuery{
		ID:      id,
		Deleted: false,
	})
}

func (svc service) GetByName(ctx context.Context, name string) (*depto.Item, error) {
	// only retrieve items that has not deleted.
	return svc.store.Get(ctx, depto.ItemQuery{
		Name:    name,
		Deleted: false,
	})
}

func (svc service) Select(ctx context.Context) ([]*depto.Item, error) {
	// only retrieve items that has not deleted.
	return svc.store.Select(ctx, depto.ItemQuery{
		Deleted: false,
	})
}

func (svc service) Create(ctx context.Context, item *depto.Item) error {
	return svc.store.Create(ctx, item)
}

func (svc service) Update(ctx context.Context, id string, item *depto.Item) (*depto.Item, error) {
	olditem, err := svc.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	olditem.Name = item.Name
	olditem.Description = item.Description

	if err := svc.store.Update(ctx, olditem); err != nil {
		return nil, err
	}

	return olditem, nil
}

func (svc service) Delete(ctx context.Context, item *depto.Item) error {
	return svc.store.Delete(ctx, item)
}
