package mongo

import (
	"context"
	"time"
)

var PrintFn func(c context.Context, method, db, coll, sql string, d time.Duration, res interface{}, err error)

func slog(c context.Context, method, db, coll, sql string, d time.Duration, res interface{}, err error) {
	if PrintFn != nil {
		PrintFn(c, method, db, coll, sql, d, res, err)
	}
}
