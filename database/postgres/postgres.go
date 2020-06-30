package postgres

import "github.com/jmoiron/sqlx"

type SQLExecutor interface {
	sqlx.QueryerContext
	sqlx.ExecerContext
}
