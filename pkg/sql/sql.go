package sql

import (
	"database/sql"
	"fmt"
)

type Client struct {
	DB *sql.DB
}

func (c Client) Insert(query string, args ...interface{}) (*int64, error) {
	stmt, err := c.DB.Prepare(query)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	rows, err := stmt.Exec(args...)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Println(rows.RowsAffected())
	lastid, err := rows.LastInsertId()
	if err != nil {
		return nil, err
	}
	return &lastid, nil
}

func (c Client) Select(query string, args ...interface{}) (*sql.Rows, error) {
	stmt, err := c.DB.Prepare(query)
	if err != nil {
		return nil, err
	}
	res, err := stmt.Query(args...)
	if err != nil {
		return nil, err
	}
	return res, err
}
