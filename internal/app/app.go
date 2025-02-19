package app

import (
	"context"
	"github.com/khostya/zero/internal/config"
	"github.com/khostya/zero/internal/http"
	"github.com/khostya/zero/internal/usecase"
	"github.com/khostya/zero/pkg/postgres"
)

func Run(ctx context.Context, cfg config.Config) error {
	pool, err := postgres.NewPostgres(ctx, cfg.PG)
	if err != nil {
		return err
	}
	defer pool.Close()

	deps := newDependencies(pool)
	useCases := usecase.NewUseCases(deps)

	return http.Run(
		ctx,
		cfg.HTTP,
		http.UseCases{
			News: useCases.Song,
		},
	)
}
