package gopostgres

import (
	"context"
	"github.com/jackc/pgx/v4"
	"io/ioutil"
)

type query struct {
	Statement string
	Args []interface{}
}

func (db Postgres) execQuery(q query) ([][]interface{}, error){
	var rows pgx.Rows
	var err error
	ctx := context.Background()
	if len(q.Args) > 0 && q.Args != nil {
		rows, err = db.connection.Query(ctx, q.Statement, q.Args...)
	} else {
		rows, err = db.connection.Query(ctx, q.Statement)
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var res [][]interface{}
	for rows.Next() {
		values, err := rows.Values()
		if err != nil {
			return nil, err
		}
		res = append(res, values)
	}

	// Any errors encountered by rows.Next or rows.Scan will be returned here
	if rows.Err() != nil {
		return nil, err
	}
	return res, nil
}

func (db Postgres) execQueriesWithTransactions(queries []query) error{
	ctx := context.Background()
	tx, err := db.connection.Begin(ctx)
	if err != nil {
		return err
	}
	// Rollback is safe to call even if the tx is already closed, so if
	// the tx commits successfully, this is a no-op
	defer tx.Rollback(ctx)

	for _, q := range queries {
		var err error
		if q.Args != nil && len(q.Args) > 0 {
			_, err = tx.Exec(ctx, q.Statement, q.Args...)
		} else {
			_, err = tx.Exec(ctx, q.Statement)
		}
		if err != nil {
			return err
		}
	}

	err = tx.Commit(ctx)
	if err != nil {
		return err
	}

	return nil
}

func getQueryFromSQLFile(path string) (*query, error) {
	buff, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	sql := string(buff)
	query := query{
		Statement: sql,
	}

	return &query, nil
}