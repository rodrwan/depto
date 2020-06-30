package repositions

import (
	"context"

	"github.com/baggyapp/depto"
	"github.com/baggyapp/depto/database"
)

type Service struct {
	Store database.RepositionsStore
}

func (svc Service) GetByID(ctx context.Context, id string) (*depto.Repositions, error) {
	// only retrieve items that has not deleted.
	return svc.Store.Get(ctx, depto.RepositionsQuery{
		ID:      id,
		Deleted: false,
	})
}

func (svc Service) GetByCollectionID(ctx context.Context, collectionID string) (*depto.Repositions, error) {
	// only retrieve items that has not deleted.
	return svc.Store.Get(ctx, depto.RepositionsQuery{
		CollectionID: collectionID,
		Deleted:      false,
	})
}

func (svc Service) Create(ctx context.Context, r *depto.Repositions) error {
	return svc.Store.Create(ctx, r)
}

func (svc Service) Update(ctx context.Context, r *depto.Repositions) error {
	return svc.Store.Update(ctx, r)
}

func (svc Service) Delete(ctx context.Context, r *depto.Repositions) error {
	return svc.Store.Delete(ctx, r)
}
