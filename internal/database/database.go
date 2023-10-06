package database

import (
	"context"
	"database/sql"
	"errors"
)

type txKey struct{}
type dbKey struct{}

func Tx(ctx context.Context) (*sql.Tx, error) {
	tx, ok := ctx.Value(txKey{}).(*sql.Tx)

	if !ok {
		return nil, errors.New("could not create a transaction")
	}

	return tx, nil
}

func getContextDB(ctx context.Context) (*sql.DB, error) {
	tx, ok := ctx.Value(dbKey{}).(*sql.DB)

	if !ok {
		return nil, errors.New("could not get database connection")
	}

	return tx, nil
}

func SetContextDB(ctx context.Context, db *sql.DB) context.Context {
	return context.WithValue(ctx, dbKey{}, db)
}

func setContextTx(ctx context.Context, tx *sql.Tx) context.Context {
	return context.WithValue(ctx, txKey{}, tx)
}

func BeginTx(ctx context.Context, fn func(context.Context) error) error {

	var db *sql.DB
	var tx *sql.Tx
	var err error

	if db, err = getContextDB(ctx); err != nil {
		return err
	}

	if tx, err = db.Begin(); err != nil {
		return err
	}

	ctx = setContextTx(ctx, tx)

	if err := fn(ctx); err != nil {
		if err := tx.Rollback(); err != nil {
			return err
		}

		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}
