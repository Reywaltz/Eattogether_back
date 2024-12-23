package database

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	pgconn "github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type DB struct {
	Conn *pgxpool.Pool
}

type PgxInterface interface {
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	Exec(ctx context.Context, sql string, args ...interface{}) (pgconn.CommandTag, error)
	Begin(ctx context.Context) (pgx.Tx, error)
}

func CreateConnection(connectionURI string) (*DB, error) {
	conn, err := pgxpool.New(context.Background(), connectionURI)

	if err != nil {
		return &DB{}, errors.New("failed to connect to db")
	}

	return &DB{Conn: conn}, nil
}

func (d *DB) Exec(ctx context.Context, query string, args ...interface{}) (pgconn.CommandTag, error) {
	tag, err := d.Conn.Exec(ctx, query, args...)
	if err != nil {
		return pgconn.CommandTag{}, err
	}

	return tag, nil

}

func (d *DB) QueryRow(ctx context.Context, query string, args ...interface{}) pgx.Row {
	return d.Conn.QueryRow(ctx, query, args...)
}

func (d *DB) Query(ctx context.Context, query string, args ...interface{}) (pgx.Rows, error) {
	rows, err := d.Conn.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	return rows, nil
}

func (d *DB) Begin(ctx context.Context) (pgx.Tx, error) {
	transaction, err := d.Conn.Begin(ctx)
	if err != nil {
		fmt.Printf("Failed to begin transaction: %v\n", err)
		return nil, err
	}

	return transaction, nil
}
