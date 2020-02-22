package gopostgres

import (
	"github.com/jackc/pgx/v4"
	"testing"
)

func TestInitDB(t *testing.T) {
	InitDB("test", "user", "pass", "localhost", nil, pgx.LogLevelDebug)
	if DB.connection == nil {
		t.Error("database could not be initialized correctly")
	}
}