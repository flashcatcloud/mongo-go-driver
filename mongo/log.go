package mongo

import "context"

var PrintFn func(c context.Context, method, db, coll, sql string, milli int64, affectRows int, err error)

func slog(c context.Context, method, db, coll, sql string, milli int64, affectRows int, err error) {
	if PrintFn != nil {
		PrintFn(c, method, db, coll, sql, milli, affectRows, err)
	}
}
