package usecase

import (
	"context"
	"github.com/khostya/zero/internal/repo"
)

type (
	transactionManager interface {
		RunRepeatableRead(ctx context.Context, fx func(ctxTX context.Context) error) error
		Unwrap(err error) error
	}

	Dependencies struct {
		Pg         repo.Repositories
		Transactor transactionManager
	}

	UseCases struct {
		Deps Dependencies
		Song Song
	}
)

func NewUseCases(deps Dependencies) UseCases {
	pg := deps.Pg

	return UseCases{
		Deps: deps,
		Song: NewNewsUseCase(NewsDeps{
			newsRepo:         pg.News,
			newsCategoryRepo: pg.NewsCategory,
			Tm:               deps.Transactor,
		}),
	}
}
