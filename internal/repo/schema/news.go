package schema

import (
	"github.com/khostya/zero/internal/domain"
	"github.com/khostya/zero/internal/repo/schema/pgmodel"
)

func NewNews(news *domain.News) *pgmodel.News {
	if news == nil {
		return nil
	}

	return &pgmodel.News{
		ID:      news.ID,
		Title:   news.Title,
		Content: news.Content,
	}
}

func NewDomainNews(song *pgmodel.News) *domain.News {
	if song == nil {
		return nil
	}
	return &domain.News{
		ID:      song.ID,
		Title:   song.Title,
		Content: song.Content,
	}
}

func NewDomainNewsWithCategory(song *pgmodel.News, category *pgmodel.NewsCategory) *domain.News {
	if song == nil {
		return nil
	}

	categories := make([]int32, 0)
	if category != nil {
		categories = append(categories, category.CategoryID)
	}

	return &domain.News{
		ID:         song.ID,
		Title:      song.Title,
		Content:    song.Content,
		Categories: categories,
	}
}

func NewDomainListNewsWithCategory(news []pgmodel.News, category []*pgmodel.NewsCategory) []*domain.News {
	if news == nil {
		return nil
	}
	if len(news) != len(category) {
		return nil
	}

	var newsIdToIndex = make(map[int32]int32)

	var res []*domain.News
	for i, n := range news {
		index, ok := newsIdToIndex[n.ID]
		if !ok {
			res = append(res, NewDomainNewsWithCategory(&n, category[i]))
			newsIdToIndex[n.ID] = int32(len(res) - 1)
		} else {
			res[index].Categories = append(res[index].Categories, category[i].CategoryID)
		}
	}

	return res
}
