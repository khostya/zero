//go:generate mockgen -source ./mocks/song.go -destination=./mocks/song_mock.go -package=mock_usecase
package usecase

import (
	"context"
	"errors"
	"github.com/khostya/zero/internal/domain"
	"github.com/khostya/zero/internal/dto"
	"github.com/khostya/zero/internal/repo/repoerr"
)

type (
	newsRepo interface {
		Save(ctx context.Context, news *domain.News) error
		GetList(ctx context.Context, param dto.GetNewsParam) ([]*domain.News, error)
	}

	newsCategoryRepo interface {
		Create(ctx context.Context, newsCategory *domain.NewsCategory) error
		Delete(ctx context.Context, newsCategory *domain.NewsCategory) error
		DeleteByNews(ctx context.Context, newsID int32) error
		CreateMulti(ctx context.Context, newsCategory []*domain.NewsCategory) error
	}

	NewsDeps struct {
		newsRepo         newsRepo
		newsCategoryRepo newsCategoryRepo
		Tm               transactionManager
	}

	Song struct {
		newsRepo         newsRepo
		newsCategoryRepo newsCategoryRepo
		tm               transactionManager
	}
)

func (uc Song) Save(ctx context.Context, param *domain.News) error {
	return uc.tm.Unwrap(uc.tm.RunRepeatableRead(ctx, func(ctx context.Context) error {
		err := uc.newsCategoryRepo.DeleteByNews(ctx, param.ID)
		if err != nil && !errors.Is(err, repoerr.ErrNotFound) {
			return err
		}

		err = uc.newsRepo.Save(ctx, param)
		if err != nil {
			return err
		}

		var newsCategory []*domain.NewsCategory
		for _, category := range param.Categories {
			newsCategory = append(newsCategory, &domain.NewsCategory{CategoryID: category, NewsID: param.ID})
		}

		err = uc.newsCategoryRepo.CreateMulti(ctx, newsCategory)
		return err
	}))
}

func (uc Song) Get(ctx context.Context, param dto.GetNewsParam) ([]*domain.News, error) {
	return uc.newsRepo.GetList(ctx, param)
}

func NewNewsUseCase(deps NewsDeps) Song {
	return Song{
		newsRepo:         deps.newsRepo,
		newsCategoryRepo: deps.newsCategoryRepo,
		tm:               deps.Tm,
	}
}
