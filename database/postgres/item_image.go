package postgres

import (
	"context"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/baggyapp/depto"
	"github.com/baggyapp/depto/database"
)

func NewItemImageStore(db SQLExecutor) database.ItemImageStore {
	return itemImageStore{
		db: db,
	}
}

type itemImageStore struct {
	db SQLExecutor
}

func (it itemImageStore) Get(ctx context.Context, isq depto.ItemImageQuery) (*depto.ItemImage, error) {
	q := squirrel.Select("*").From("item_images").Where("deleted_at is null")

	if isq.ID > 0 {
		q = q.Where("id = ?", isq.ID)
	}

	if isq.ItemID > 0 {
		q = q.Where("item_info_id = ?", isq.ItemID)
	}

	if !isq.Deleted {
		q = q.Where("deleted_at is null")
	}

	sqlString, args, err := q.PlaceholderFormat(squirrel.Dollar).ToSql()
	if err != nil {
		return nil, err
	}

	var i *depto.ItemImage

	row := it.db.QueryRowxContext(ctx, sqlString, args...)
	if err := row.StructScan(i); err != nil {
		return nil, err
	}

	return i, nil
}

func (it itemImageStore) Select(ctx context.Context, isq depto.ItemImageQuery) ([]*depto.ItemImage, error) {
	q := squirrel.Select("*").From("item_images")

	if isq.ItemID > 0 {
		q = q.Where("item_info_id = ?", isq.ItemID)
	}

	if !isq.Deleted {
		q = q.Where("deleted_at is null")
	}

	sqlString, args, err := q.PlaceholderFormat(squirrel.Dollar).ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := it.db.QueryxContext(ctx, sqlString, args...)
	if err != nil {
		return nil, err
	}

	items := []*depto.ItemImage{}

	for rows.Next() {
		var item depto.ItemImage
		if err := rows.StructScan(&item); err != nil {
			return nil, err
		}

		items = append(items, &item)
	}

	return items, nil
}

func (it itemImageStore) Create(ctx context.Context, i *depto.ItemImage) error {
	sqlString, args, err := squirrel.Insert("item_images").
		Columns("item_info_id", "image_url").
		Values(i.ItemInfoID, i.ImageURL).
		Suffix("returning *").
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return err
	}

	row := it.db.QueryRowxContext(ctx, sqlString, args...)
	if err := row.StructScan(i); err != nil {
		return err
	}

	return nil
}

func (it itemImageStore) Update(ctx context.Context, i *depto.ItemImage) error {
	sqlString, args, err := squirrel.Update("item_images").
		Set("item_info_id", i.ItemInfoID).
		Set("image_url", i.ImageURL).
		Where("id = ?", i.ID).
		Suffix("returning *").
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return err
	}

	row := it.db.QueryRowxContext(ctx, sqlString, args...)
	if err := row.StructScan(i); err != nil {
		return err
	}

	return nil
}

func (it itemImageStore) Delete(ctx context.Context, i *depto.ItemImage) error {
	sqlString, args, err := squirrel.Update("item_images").
		Set("deleted_at", time.Now()).
		Where("id = ?", i.ID).
		Suffix("returning *").
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return err
	}

	row := it.db.QueryRowxContext(ctx, sqlString, args...)
	if err := row.StructScan(i); err != nil {
		return err
	}

	return nil
}
