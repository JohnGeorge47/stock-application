package sql

import "database/sql"

type Isql interface {
	Insert(query string, args ...interface{}) (*int64, error)
	Select(query string, args ...interface{}) (*sql.Rows, error)
}
