package postgres

import (
	"context"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/baggyapp/depto"
)

type InventoryStore struct {
	DB SQLExecutor
}

func (is InventoryStore) Get(ctx context.Context, sq depto.InventoryQuery) (*depto.Inventory, error) {
	q := squirrel.Select("*").From("inventory").Where("deleted_at is null")

	if sq.ID != "" {
		q = q.Where("id = ?", sq.ID)
	}

	if !sq.Deleted {
		q = q.Where("deleted_at is null")
	}

	sqlString, args, err := q.PlaceholderFormat(squirrel.Dollar).ToSql()
	if err != nil {
		return nil, err
	}

	var s *depto.Inventory

	row := is.DB.QueryRowxContext(ctx, sqlString, args...)
	if err := row.StructScan(s); err != nil {
		return nil, err
	}

	return s, nil
}

func (is InventoryStore) Select(ctx context.Context, sq depto.InventoryQuery) ([]*depto.Inventory, error) {
	q := squirrel.Select("*").From("inventory")

	if !sq.Deleted {
		q = q.Where("deleted_at is null")
	}

	sqlString, args, err := q.PlaceholderFormat(squirrel.Dollar).ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := is.DB.QueryxContext(ctx, sqlString, args...)
	if err != nil {
		return nil, err
	}

	inventory := []*depto.Inventory{}

	for rows.Next() {
		var stock depto.Inventory
		if err := rows.StructScan(&stock); err != nil {
			return nil, err
		}

		inventory = append(inventory, &stock)
	}

	return inventory, nil
}

func (is InventoryStore) Create(ctx context.Context, i *depto.Inventory) error {
	sqlString, args, err := squirrel.Insert("inventory").
		Columns(
			"name",
		).
		Values(
			i.Name,
		).
		Suffix("returning *").
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return err
	}

	row := is.DB.QueryRowxContext(ctx, sqlString, args...)
	if err := row.StructScan(i); err != nil {
		return err
	}

	return nil
}

func (is InventoryStore) Update(ctx context.Context, i *depto.Inventory) error {
	sqlString, args, err := squirrel.Update("inventory").
		Set("name", i.Name).
		Where("id = ?", i.ID).
		Suffix("returning *").
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return err
	}

	row := is.DB.QueryRowxContext(ctx, sqlString, args...)
	if err := row.StructScan(i); err != nil {
		return err
	}

	return nil
}

func (is InventoryStore) Delete(ctx context.Context, i *depto.Inventory) error {
	sqlString, args, err := squirrel.Update("inventory").
		Set("deleted_at", time.Now()).
		Where("id = ?", i.ID).
		Suffix("returning *").
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return err
	}

	row := is.DB.QueryRowxContext(ctx, sqlString, args...)
	if err := row.StructScan(i); err != nil {
		return err
	}

	return nil
}
