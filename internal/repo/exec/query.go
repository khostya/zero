package exec

import (
	"context"
	"database/sql"
	sq "github.com/Masterminds/squirrel"
	"gopkg.in/reform.v1"
)

func Query(ctx context.Context, query sq.Sqlizer, db *reform.Querier) (*sql.Rows, error) {
	rawQuery, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := db.QueryContext(ctx, rawQuery, args...)
	if err != nil {
		return nil, err
	}

	return rows, nil
}
