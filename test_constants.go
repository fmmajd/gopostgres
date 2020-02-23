package gopostgres

var testTable = "testing"

var createTestTable = "CREATE TABLE IF NOT EXISTS "+testTable+
	"(id bigserial primary key," +
	"username varchar not null unique," +
	"last_updated timestamp default now());"

var truncateTestTable = "truncate table testing"

var insertOneRowInTestTable = "insert into "+testTable+" (username, last_updated) values ($1, $2)"