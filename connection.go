package gopostgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
)

func newConnection (db string, user string, password string, host string, logger pgx.Logger, logLevel pgx.LogLevel) (*pgx.Conn, error){
	configStr := fmt.Sprintf("user=%s password=%s host=%s port=5432 dbname=%s statement_cache_capacity=0", user, password, host, db)
	config, err := pgx.ParseConfig(configStr)
	if err != nil {
		return nil, err
	}
	config.Logger = logger
	config.LogLevel = logLevel

	conn, err := pgx.ConnectConfig(context.Background(), config)

	return conn, err
}

