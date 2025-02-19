package postgres

import (
	"context"
	"database/sql"
	"github.com/khostya/zero/internal/config"
	_ "github.com/lib/pq"
	"gopkg.in/reform.v1"
	"gopkg.in/reform.v1/dialects/postgresql"
	"log"
)

type Postgres struct {
	db  *reform.DB
	sql *sql.DB
}

func NewPostgres(ctx context.Context, cfg config.PG) (*Postgres, error) {
	sqlDB, err := sql.Open("postgres", cfg.URL)
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxOpenConns(cfg.MaxOpenConns)
	sqlDB.SetMaxIdleConns(cfg.MaxIdleConns)
	sqlDB.SetConnMaxIdleTime(cfg.ConnMaxIdleTime)
	sqlDB.SetConnMaxLifetime(cfg.ConnMaxLifetime)

	db := reform.NewDB(sqlDB, postgresql.Dialect, reform.NewPrintfLogger(log.Printf))
	return &Postgres{db: db, sql: sqlDB}, nil
}

func (p *Postgres) GetDB() *reform.DB {
	return p.db
}

func (p *Postgres) Close() error {
	return p.sql.Close()
}
