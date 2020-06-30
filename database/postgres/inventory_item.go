package postgres

import (
	"context"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/baggyapp/depto"
)

type InventoryItemStore struct {
	DB SQLExecutor
}

func (is InventoryItemStore) Get(ctx context.Context, sq depto.InventoryItemQuery) (*depto.InventoryItem, error) {
	q := squirrel.Select("*").From("inventory_items").Where("deleted_at is null")

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

	var ii depto.InventoryItem

	row := is.DB.QueryRowxContext(ctx, sqlString, args...)
	if err := row.StructScan(&ii); err != nil {
		return nil, err
	}

	return &ii, nil
}

func (is InventoryItemStore) Select(ctx context.Context, sq depto.InventoryItemQuery) ([]*depto.InventoryItem, error) {
	q := squirrel.Select("*").From("inventory")

	if sq.InventoryID != "" {
		q = q.Where("inventory_id = ?", sq.InventoryID)
	}

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

	inventory := []*depto.InventoryItem{}

	for rows.Next() {
		var stock depto.InventoryItem
		if err := rows.StructScan(&stock); err != nil {
			return nil, err
		}

		inventory = append(inventory, &stock)
	}

	return inventory, nil
}

func (is InventoryItemStore) Create(ctx context.Context, ii *depto.InventoryItem) error {
	sqlString, args, err := squirrel.Insert("inventory_items").
		Columns(
			"inventory_id",
			"item_id",
			"count",
			"last_use",
		).
		Values(
			ii.InventoryID,
			ii.ItemID,
			ii.Count,
			ii.LastUse,
		).
		Suffix("returning *").
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return err
	}

	row := is.DB.QueryRowxContext(ctx, sqlString, args...)
	if err := row.StructScan(ii); err != nil {
		return err
	}

	return nil
}

func (is InventoryItemStore) Update(ctx context.Context, ii *depto.InventoryItem) error {
	sqlString, args, err := squirrel.Update("inventory_items").
		Set("inventory_id", ii.InventoryID).
		Set("item_id", ii.ItemID).
		Set("count", ii.Count).
		Set("last_use", ii.LastUse).
		Where("id = ?", ii.ID).
		Suffix("returning *").
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return err
	}

	row := is.DB.QueryRowxContext(ctx, sqlString, args...)
	if err := row.StructScan(ii); err != nil {
		return err
	}

	return nil
}

func (is InventoryItemStore) Delete(ctx context.Context, ii *depto.InventoryItem) error {
	sqlString, args, err := squirrel.Update("inventory_items").
		Set("deleted_at", time.Now()).
		Where("id = ?", ii.ID).
		Suffix("returning *").
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return err
	}

	row := is.DB.QueryRowxContext(ctx, sqlString, args...)
	if err := row.StructScan(ii); err != nil {
		return err
	}

	return nil
}
