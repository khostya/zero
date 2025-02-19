package schema

import (
	"github.com/khostya/zero/internal/domain"
	"github.com/khostya/zero/internal/repo/schema/pgmodel"
)

func NewNewsCategory(newsCategory *domain.NewsCategory) *pgmodel.NewsCategory {
	if newsCategory == nil {
		return nil
	}

	return &pgmodel.NewsCategory{
		NewsID:     newsCategory.NewsID,
		CategoryID: newsCategory.CategoryID,
	}
}
