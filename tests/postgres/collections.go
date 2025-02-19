//go:build integration

package postgres

import (
	"github.com/brianvoe/gofakeit/v7"
	"github.com/khostya/zero/internal/domain"
)

func NewNews() *domain.News {
	return &domain.News{
		ID:         gofakeit.Int32(),
		Title:      "title",
		Content:    "content",
		Categories: nil,
	}
}

func NewNewsCategory(newsID int32) *domain.NewsCategory {
	return &domain.NewsCategory{
		NewsID:     newsID,
		CategoryID: gofakeit.Int32(),
	}
}
func NewNewsCategories(newsID int32) []*domain.NewsCategory {
	return []*domain.NewsCategory{
		NewNewsCategory(newsID),
		NewNewsCategory(newsID),
	}
}
