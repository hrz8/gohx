package main

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5"
)

func connect(ctx context.Context) *pgx.Conn {
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		os.Exit(0)
		return nil
	}

	conn, err := pgx.Connect(ctx, databaseURL)
	if err != nil {
		panic(err)
	}

	return conn
}
