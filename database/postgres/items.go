package postgres

import (
	"context"
	"errors"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/baggyapp/depto"
	"github.com/baggyapp/depto/database"
)

func NewItemsStore(db SQLExecutor) database.ItemStore {
	return itemsStore{
		db: db,
	}
}

type itemsStore struct {
	db SQLExecutor
}

func (it itemsStore) Get(ctx context.Context, isq depto.ItemQuery) (*depto.Item, error) {
	if isq.ID == "" && isq.Name == "" {
		return nil, errors.New("must provide an id or a name")
	}

	q := squirrel.Select("*").From("items")

	if isq.ID != "" {
		q = q.Where("id = ?", isq.ID)
	}

	if isq.Name != "" {
		q = q.Where("name = ?", isq.Name)
	}

	if !isq.Deleted {
		q = q.Where("deleted_at is null")
	}

	sqlString, args, err := q.PlaceholderFormat(squirrel.Dollar).ToSql()
	if err != nil {
		return nil, err
	}

	var i depto.Item

	row := it.db.QueryRowxContext(ctx, sqlString, args...)
	if err := row.StructScan(&i); err != nil {
		return nil, err
	}

	return &i, nil
}

func (it itemsStore) Select(ctx context.Context, isq depto.ItemQuery) ([]*depto.Item, error) {
	q := squirrel.Select("*").From("items")

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

	items := []*depto.Item{}

	for rows.Next() {
		var item depto.Item
		if err := rows.StructScan(&item); err != nil {
			return nil, err
		}

		items = append(items, &item)
	}

	return items, nil
}

func (it itemsStore) Create(ctx context.Context, i *depto.Item) error {
	sqlString, args, err := squirrel.Insert("items").
		Columns("name", "description").
		Values(i.Name, i.Description).
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

func (it itemsStore) Update(ctx context.Context, i *depto.Item) error {
	sqlString, args, err := squirrel.Update("items").
		Set("name", i.Name).
		Set("description", i.Description).
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

func (it itemsStore) Delete(ctx context.Context, i *depto.Item) error {
	sqlString, args, err := squirrel.Update("items").
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
