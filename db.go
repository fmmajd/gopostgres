package gopostgres

import (
	"github.com/jackc/pgx/v4"
	"log"
	"sync"
)

type Postgres struct {
	connection *pgx.Conn
}

var DB Postgres
var once sync.Once

func InitDB(db string, user string, password string, host string, logger pgx.Logger, logLevel pgx.LogLevel) {
	once.Do(func() {
		con, err := newConnection(db, user, password, host, logger, logLevel)
		if err != nil {
			log.Fatalln(err)
		}
		DB = Postgres{
			connection: con,
		}
	})
}
