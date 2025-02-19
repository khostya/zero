//go:build integration

package postgres

import (
	"context"
	"github.com/khostya/zero/tests/postgres/postgresql"
	"os"
	"testing"
)

var (
	db *postgresql.DBPool
)

const (
	newsTable         = "news"
	newsCategoryTable = "news_category"
)

func TestMain(m *testing.M) {
	db = postgresql.NewFromEnv()

	code := m.Run()
	truncate()
	db.Close()

	os.Exit(code)
}

func truncate() {
	db.TruncateTable(context.Background(), newsCategoryTable, newsTable)
}
