package client

import (
	"context"
	"fmt"

	"github.com/gmcriptobox/otus-go-final-project/internal/config"
	"github.com/jmoiron/sqlx"
)

type PostgresSQL struct {
	DB     *sqlx.DB
	config config.Config
}

func NewPostgresSQL(config config.Config) *PostgresSQL {
	return &PostgresSQL{
		config: config,
	}
}

func (p *PostgresSQL) Connect(ctx context.Context) error {
	db, err := sqlx.Open("postgres", p.config.PSQL.DSN)
	if err != nil {
		return fmt.Errorf("cannot open pgx driver: %w", err)
	}

	p.DB = db
	return p.DB.PingContext(ctx)
}

func (p *PostgresSQL) Close() error {
	return p.DB.Close()
}
