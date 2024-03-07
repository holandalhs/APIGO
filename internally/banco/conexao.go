package banco

import (
	"context"
	"fmt"

	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

var Conn *pgxpool.Pool

func NewConnection(connectionString string) (*pgxpool.Pool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	pool, err := pgxpool.Connect(ctx, connectionString)
	if err != nil {
		return nil, fmt.Errorf("Falha de conexão com o banco: %w", err)
	}

	Conn = pool
	return Conn, nil
}
