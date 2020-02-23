package gopostgres

var testTableForFindBy = "testing1"
var testTableForFindAllBy = "testing2"

var createTestTable = "CREATE TABLE IF NOT EXISTS %s"+
	"(id bigserial primary key," +
	"username varchar not null unique," +
	"last_updated timestamp default now());"

var truncateTestTable = "truncate table %s"

var insertOneRowInTestTable = "insert into %s (username, last_updated) values ($1, $2)"