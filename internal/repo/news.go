//go:generate mockgen -source ./mocks/group.go -destination=./mocks/group_mock.go -package=mock_repository
package repo

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/khostya/zero/internal/domain"
	"github.com/khostya/zero/internal/dto"
	"github.com/khostya/zero/internal/repo/exec"
	"github.com/khostya/zero/internal/repo/schema"
	"github.com/khostya/zero/internal/repo/schema/pgmodel"
	"github.com/khostya/zero/pkg/postgres/transactor"
)

const (
	groupsTable = "effective.groups"
)

type (
	News struct {
		queryEngineProvider transactor.QueryEngineProvider
	}
)

func (g News) Save(ctx context.Context, news *domain.News) error {
	db := g.queryEngineProvider.GetQueryEngine(ctx)

	record := schema.NewNews(news)

	err := db.Save(record)
	return err
}

func (g News) GetList(ctx context.Context, param dto.GetNewsParam) ([]*domain.News, error) {
	db := g.queryEngineProvider.GetQueryEngine(ctx)

	columns := db.QualifiedColumns(pgmodel.NewsTable)
	columns = append(columns, db.QualifiedColumns(pgmodel.NewsCategoryView)...)

	query := sq.Select(columns...).
		From("news").
		LeftJoin("news_category ON news.id = news_category.news_id")

	if param.Page != nil {
		offset, err := param.Page.Offset()
		if err != nil {
			return nil, err
		}
		query = query.
			Offset(uint64(offset)).
			Limit(uint64(param.Page.Limit()))
	}

	rows, err := exec.Query(ctx, query, db)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var (
		newsList       []pgmodel.News
		newsCategories []*pgmodel.NewsCategory
	)

	for rows.Next() {
		var news pgmodel.News
		var newsCategory *pgmodel.NewsCategory

		var (
			newsID     *int32
			categoryID *int32
		)

		pointers := news.Pointers()
		pointers = append(pointers, &newsID, &categoryID)

		if err = rows.Scan(pointers...); err != nil {
			return nil, err
		}

		newsList = append(newsList, news)
		if newsID != nil && categoryID != nil {
			newsCategory = &pgmodel.NewsCategory{
				NewsID:     *newsID,
				CategoryID: *categoryID,
			}
		}

		newsCategories = append(newsCategories, newsCategory)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return schema.NewDomainListNewsWithCategory(newsList, newsCategories), nil
}

func NewNewsRepo(provider transactor.QueryEngineProvider) *News {
	return &News{provider}
}
