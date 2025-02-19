package http

import (
	"context"
	"github.com/khostya/zero/internal/domain"
	"github.com/khostya/zero/internal/dto"
	"github.com/khostya/zero/internal/usecase"
	"github.com/khostya/zero/pkg/validator"
)

var (
	_ newsUseCase = (*usecase.Song)(nil)
)

type (
	newsUseCase interface {
		Save(ctx context.Context, param *domain.News) error
		Get(ctx context.Context, param dto.GetNewsParam) ([]*domain.News, error)
	}

	UseCases struct {
		News newsUseCase
	}

	server struct {
		useCases  UseCases
		validator *validator.Validator
	}
)

func newServer(useCases UseCases) (*server, error) {
	validator, err := validator.NewValidate()
	if err != nil {
		return nil, err
	}

	return &server{
		useCases:  useCases,
		validator: validator,
	}, nil
}
