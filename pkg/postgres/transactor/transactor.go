package transactor

import (
	"context"
	"database/sql"
	"errors"
	"github.com/khostya/zero/pkg/postgres"
	"gopkg.in/reform.v1"
)

const key = "tx"

type (
	// Transactor .
	Transactor interface {
		RunRepeatableRead(ctx context.Context, fx func(ctxTX context.Context) error) error
		Unwrap(err error) error
		GetQueryEngine(ctx context.Context) *reform.Querier
	}

	QueryEngineProvider interface {
		GetQueryEngine(ctx context.Context) *reform.Querier // tx OR pool
	}

	TransactionManager struct {
		pool *reform.DB
	}
)

func NewTransactionManager(pool *postgres.Postgres) *TransactionManager {
	return &TransactionManager{pool.GetDB()}
}

func (tm TransactionManager) RunRepeatableRead(ctx context.Context, fx func(ctx context.Context) error) error {
	tx, err := tm.pool.BeginTx(ctx, &sql.TxOptions{
		Isolation: sql.LevelRepeatableRead,
	})

	if err != nil {
		return TransactionError{Inner: err}
	}
	if err := fx(context.WithValue(ctx, key, tx)); err != nil {
		return TransactionError{Inner: err, Rollback: tx.Rollback()}
	}

	if err := tx.Commit(); err != nil {
		return TransactionError{Inner: err, Rollback: tx.Rollback()}
	}

	return nil
}

func (tm TransactionManager) Unwrap(err error) error {
	if err == nil {
		return nil
	}

	var transactionError TransactionError
	ok := errors.As(err, &transactionError)
	if !ok {
		return err
	}
	return transactionError.Inner
}

func (tm TransactionManager) GetQueryEngine(ctx context.Context) *reform.Querier {
	tx, ok := ctx.Value(key).(*reform.Querier)

	if ok && tx != nil {
		return tx
	}

	return tm.pool.Querier.WithContext(ctx)
}
