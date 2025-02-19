package app

import (
	"github.com/khostya/zero/internal/repo"
	"github.com/khostya/zero/internal/usecase"
	"github.com/khostya/zero/pkg/postgres"
	"github.com/khostya/zero/pkg/postgres/transactor"
)

func newDependencies(pool *postgres.Postgres) usecase.Dependencies {
	transactor := transactor.NewTransactionManager(pool)
	pgRepositories := repo.NewRepositories(transactor)

	return usecase.Dependencies{
		Pg:         pgRepositories,
		Transactor: transactor,
	}
}
