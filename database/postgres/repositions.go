package postgres

import (
	"context"
	"errors"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/baggyapp/depto"
)

type RepositionsStore struct {
	DB SQLExecutor
}

func (it RepositionsStore) Get(ctx context.Context, isq depto.RepositionsQuery) (*depto.Repositions, error) {
	q := squirrel.Select("*").From("items").Where("deleted_at is null")

	if isq.CollectionID == "" {
		return nil, errors.New("must provide an collection_id")
	}

	if isq.CollectionID != "" {
		q = q.Where("collection_id = ?", isq.CollectionID)
	}

	if !isq.Deleted {
		q = q.Where("deleted_at is null")
	}

	sqlString, args, err := q.PlaceholderFormat(squirrel.Dollar).ToSql()
	if err != nil {
		return nil, err
	}

	var r *depto.Repositions

	row := it.DB.QueryRowxContext(ctx, sqlString, args...)
	if err := row.StructScan(r); err != nil {
		return nil, err
	}

	return r, nil
}

func (it RepositionsStore) Select(ctx context.Context, isq depto.RepositionsQuery) ([]*depto.Repositions, error) {
	q := squirrel.Select("*").From("items")

	if !isq.Deleted {
		q = q.Where("deleted_at is null")
	}

	sqlString, args, err := q.PlaceholderFormat(squirrel.Dollar).ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := it.DB.QueryxContext(ctx, sqlString, args...)
	if err != nil {
		return nil, err
	}

	repos := []*depto.Repositions{}

	for rows.Next() {
		var rep depto.Repositions
		if err := rows.StructScan(&rep); err != nil {
			return nil, err
		}

		repos = append(repos, &rep)
	}

	return repos, nil
}

func (it RepositionsStore) Create(ctx context.Context, r *depto.Repositions) error {
	sqlString, args, err := squirrel.Insert("items").
		Columns("collection_id").
		Values(r.CollectionID).
		Suffix("returning *").
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return err
	}

	row := it.DB.QueryRowxContext(ctx, sqlString, args...)
	if err := row.StructScan(r); err != nil {
		return err
	}

	return nil
}

func (it RepositionsStore) Update(ctx context.Context, r *depto.Repositions) error {
	sqlString, args, err := squirrel.Update("repositions").
		Set("collection_id", r.CollectionID).
		Where("id = ?", r.ID).
		Suffix("returning *").
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return err
	}

	row := it.DB.QueryRowxContext(ctx, sqlString, args...)
	if err := row.StructScan(r); err != nil {
		return err
	}

	return nil
}

func (it RepositionsStore) Delete(ctx context.Context, r *depto.Repositions) error {
	sqlString, args, err := squirrel.Update("repositions").
		Set("deleted_at", time.Now()).
		Where("id = ?", r.ID).
		Suffix("returning *").
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return err
	}

	row := it.DB.QueryRowxContext(ctx, sqlString, args...)
	if err := row.StructScan(r); err != nil {
		return err
	}

	return nil
}
