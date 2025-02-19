package repo

import (
	"github.com/khostya/zero/pkg/postgres/transactor"
)

type Repositories struct {
	NewsCategory *NewsCategory
	News         *News
}

func NewRepositories(provider transactor.QueryEngineProvider) Repositories {
	return Repositories{
		NewsCategory: NewNewsCategoryRepo(provider),
		News:         NewNewsRepo(provider),
	}
}
