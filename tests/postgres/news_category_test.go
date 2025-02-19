//go:build integration

package postgres

import (
	"context"
	"github.com/khostya/zero/internal/domain"
	"github.com/khostya/zero/internal/repo"
	"github.com/khostya/zero/pkg/postgres/transactor"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"testing"
)

type GroupTestSuite struct {
	suite.Suite
	ctx          context.Context
	newsRepo     *repo.News
	newsCategory *repo.NewsCategory
	transactor   *transactor.TransactionManager
}

func TestGroup(t *testing.T) {
	suite.Run(t, new(GroupTestSuite))
}

func (s *GroupTestSuite) SetupSuite() {
	s.transactor = transactor.NewTransactionManager(db.GetDB())
	s.newsRepo = repo.NewNewsRepo(s.transactor)
	s.newsCategory = repo.NewNewsCategoryRepo(s.transactor)
	s.ctx = context.Background()
}

func (s *GroupTestSuite) TestCreate() {
	_ = s.create()
}

func (s *GroupTestSuite) TestDelete() {
	newsCategory := s.create()

	err := s.newsCategory.Delete(s.ctx, newsCategory)
	require.NoError(s.T(), err)
}

func (s *GroupTestSuite) create() *domain.NewsCategory {
	news := NewNews()

	err := s.newsRepo.Save(s.ctx, news)
	require.NoError(s.T(), err)

	newsCategory := NewNewsCategory(news.ID)
	err = s.newsCategory.Create(s.ctx, newsCategory)

	require.NoError(s.T(), err)

	return newsCategory
}
