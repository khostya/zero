package repo

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/khostya/zero/internal/domain"
	"github.com/khostya/zero/internal/repo/exec"
	"github.com/khostya/zero/internal/repo/schema"
	"github.com/khostya/zero/internal/repo/schema/pgmodel"
	"github.com/khostya/zero/pkg/postgres/transactor"
	"gopkg.in/reform.v1"
)

type (
	NewsCategory struct {
		queryEngineProvider transactor.QueryEngineProvider
	}
)

func (s NewsCategory) Create(ctx context.Context, newsCategory *domain.NewsCategory) error {
	db := s.queryEngineProvider.GetQueryEngine(ctx)

	record := schema.NewNewsCategory(newsCategory)

	err := db.Insert(record)
	return err
}

func (s NewsCategory) CreateMulti(ctx context.Context, newsCategory []*domain.NewsCategory) error {
	db := s.queryEngineProvider.GetQueryEngine(ctx)

	var structs []reform.Struct
	for _, record := range newsCategory {
		structs = append(structs, schema.NewNewsCategory(record))
	}

	err := db.InsertMulti(structs...)
	return err
}

func (s NewsCategory) Delete(ctx context.Context, newsCategory *domain.NewsCategory) error {
	db := s.queryEngineProvider.GetQueryEngine(ctx)

	record := schema.NewNewsCategory(newsCategory)

	query := sq.Delete(pgmodel.NewsCategoryView.Name()).
		Where("news_id = $1", record.NewsID).
		Where("category_id = $2", record.CategoryID)

	return exec.Delete(ctx, query, db)
}

func (s NewsCategory) DeleteByNews(ctx context.Context, newsID int32) error {
	db := s.queryEngineProvider.GetQueryEngine(ctx)

	query := sq.Delete(pgmodel.NewsCategoryView.Name()).
		Where("news_id = $1", newsID)

	return exec.Delete(ctx, query, db)
}

func NewNewsCategoryRepo(provider transactor.QueryEngineProvider) *NewsCategory {
	return &NewsCategory{provider}
}
