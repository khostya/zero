//go:build integration

package postgres

import (
	"context"
	"github.com/khostya/zero/internal/domain"
	"github.com/khostya/zero/internal/dto"
	"github.com/khostya/zero/internal/repo"
	"github.com/khostya/zero/pkg/postgres/transactor"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"testing"
)

type NewsTestSuite struct {
	suite.Suite
	ctx              context.Context
	newsCategoryRepo *repo.NewsCategory
	newsRepo         *repo.News
	transactor       *transactor.TransactionManager
}

func TestNews(t *testing.T) {
	suite.Run(t, new(NewsTestSuite))
}

func (s *NewsTestSuite) SetupSuite() {
	s.transactor = transactor.NewTransactionManager(db.GetDB())
	s.newsRepo = repo.NewNewsRepo(s.transactor)
	s.newsCategoryRepo = repo.NewNewsCategoryRepo(s.transactor)
	s.ctx = context.Background()
}

func (s *NewsTestSuite) TestCreate() {
	_ = s.create(nil)
}

func (s *NewsTestSuite) TestCreateWithCategories() {
	_ = s.create([]int32{1, 2, 3})
}

func (s *NewsTestSuite) TestGet() {
	truncate()

	news := s.create([]int32{1, 2, 3})

	newsList, err := s.newsRepo.GetList(s.ctx, dto.GetNewsParam{
		Page: &dto.Page{
			Page: 1,
			Size: 10,
		},
	})
	require.NoError(s.T(), err)
	require.Lenf(s.T(), newsList, 1, "expected 1 news, got %d", len(newsList))
	require.Equal(s.T(), news, newsList[0])
}

func (s *NewsTestSuite) create(categories []int32) *domain.News {
	news := NewNews()
	news.Categories = categories

	err := s.newsRepo.Save(s.ctx, news)
	require.NoError(s.T(), err)

	var res []*domain.NewsCategory
	for _, c := range categories {
		res = append(res, &domain.NewsCategory{
			NewsID:     news.ID,
			CategoryID: c,
		})
	}
	err = s.newsCategoryRepo.CreateMulti(s.ctx, res)
	require.NoError(s.T(), err)

	return news
}
