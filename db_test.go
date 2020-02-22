package gopostgres

import (
	"github.com/jackc/pgx/v4"
	"testing"
)

func TestDbNotInitializedBeforeInitCall(t *testing.T) {
	db := DB
	if db.connection != nil {
		t.Error("db is initialized without the call to InitDB function")
	}
}

func TestInitDB(t *testing.T) {
	InitDB("test", "user", "pass", "localhost", nil, pgx.LogLevelDebug)
	if DB.connection == nil {
		t.Error("database could not be initialized correctly")
	}
}