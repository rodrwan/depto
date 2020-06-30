package postgres

import (
	"context"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/baggyapp/depto"
	"github.com/baggyapp/depto/database"
)

func NewItemInfoStore(db SQLExecutor) database.ItemInfoStore {
	return itemInfoStore{
		db: db,
	}
}

type itemInfoStore struct {
	db SQLExecutor
}

func (iit itemInfoStore) Get(ctx context.Context, isq depto.ItemInfoQuery) (*depto.ItemInfo, error) {
	q := squirrel.Select("*").From("item_info").Where("deleted_at is null")

	if isq.ID > 0 {
		q = q.Where("id = ?", isq.ID)
	}

	if isq.ItemID != "" {
		q = q.Where("item_id = ?", isq.ItemID)
	}

	if !isq.Deleted {
		q = q.Where("deleted_at is null")
	}

	sqlString, args, err := q.PlaceholderFormat(squirrel.Dollar).ToSql()
	if err != nil {
		return nil, err
	}

	var ii depto.ItemInfo

	row := iit.db.QueryRowxContext(ctx, sqlString, args...)
	if err := row.StructScan(&ii); err != nil {
		return nil, err
	}

	return &ii, nil
}

func (iit itemInfoStore) Select(ctx context.Context, isq depto.ItemInfoQuery) ([]*depto.ItemInfo, error) {
	q := squirrel.Select("*").From("item_info")

	if isq.ItemID != "" {
		q = q.Where("item_id = ?", isq.ItemID)
	}

	if !isq.Deleted {
		q = q.Where("deleted_at is null")
	}

	sqlString, args, err := q.PlaceholderFormat(squirrel.Dollar).ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := iit.db.QueryxContext(ctx, sqlString, args...)
	if err != nil {
		return nil, err
	}

	items := []*depto.ItemInfo{}

	for rows.Next() {
		var item depto.ItemInfo
		if err := rows.StructScan(&item); err != nil {
			return nil, err
		}

		items = append(items, &item)
	}

	return items, nil
}

func (iit itemInfoStore) Create(ctx context.Context, i *depto.ItemInfo) error {
	sqlString, args, err := squirrel.Insert("item_info").
		Columns(
			"item_id",
			"brand",
			"price",
			"unit",
			"purchase_date",
			"purchase_place",
			"expiration_date",
		).
		Values(
			i.ItemID,
			i.Brand,
			i.Price,
			i.Unit,
			i.PurchaseDate,
			i.PurchasePlace,
			i.ExpirationDate,
		).
		Suffix("returning *").
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return err
	}

	row := iit.db.QueryRowxContext(ctx, sqlString, args...)
	if err := row.StructScan(i); err != nil {
		return err
	}

	return nil
}

func (iit itemInfoStore) Update(ctx context.Context, i *depto.ItemInfo) error {
	sqlString, args, err := squirrel.Update("item_info").
		Set("item_id", i.ItemID).
		Set("brand", i.Brand).
		Set("price", i.Price).
		Set("purchase_date", i.PurchaseDate).
		Set("purchase_place", i.PurchasePlace).
		Set("ExpirationDate", i.ExpirationDate).
		Where("id = ?", i.ID).
		Suffix("returning *").
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return err
	}

	row := iit.db.QueryRowxContext(ctx, sqlString, args...)
	if err := row.StructScan(i); err != nil {
		return err
	}

	return nil
}

func (iit itemInfoStore) Delete(ctx context.Context, i *depto.ItemInfo) error {
	sqlString, args, err := squirrel.Update("item_info").
		Set("deleted_at", time.Now()).
		Where("id = ?", i.ID).
		Suffix("returning *").
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return err
	}

	row := iit.db.QueryRowxContext(ctx, sqlString, args...)
	if err := row.StructScan(i); err != nil {
		return err
	}

	return nil
}
