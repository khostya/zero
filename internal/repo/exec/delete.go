package exec

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/khostya/zero/internal/repo/repoerr"
	"gopkg.in/reform.v1"
)

func Delete(ctx context.Context, query sq.Sqlizer, db *reform.Querier) error {
	sql, args, err := query.ToSql()
	if err != nil {
		return err
	}

	tag, err := db.ExecContext(ctx, sql, args...)
	if err != nil {
		return err
	}

	if n, _ := tag.RowsAffected(); n == 0 {
		return repoerr.ErrNotFound
	}

	return nil
}
