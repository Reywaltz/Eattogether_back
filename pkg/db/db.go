package database

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	pgconn "github.com/jackc/pgx/v5/pgconn"
)

type DB struct {
	Conn *pgx.Conn
}

type PgxInterface interface {
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Row, error)
	Exec(ctx context.Context, sql string, args ...interface{}) (pgconn.CommandTag, error)
}

func CreateConnection(connectionURI string) (*DB, error) {
	conn, err := pgx.Connect(context.Background(), connectionURI)

	if err != nil {
		return &DB{}, errors.New("failed to connect to db")
	}

	return &DB{Conn: conn}, nil
}

func (d *DB) Exec(ctx context.Context, query string, args ...interface{}) (pgconn.CommandTag, error) {
	tag, err := d.Conn.Exec(ctx, query, args)
	if err != nil {
		return pgconn.CommandTag{}, err
	}

	return tag, nil

}

func (d *DB) QueryRow(ctx context.Context, query string, args ...interface{}) pgx.Row {
	return d.Conn.QueryRow(ctx, query, args)
}

func (d *DB) Query(ctx context.Context, query string, args ...interface{}) (pgx.Row, error) {
	row, err := d.Conn.Query(ctx, query, args)
	if err != nil {
		return nil, err
	}

	return row, nil
}
