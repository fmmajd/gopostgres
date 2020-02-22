package gopostgres

var queryTableExists = "select exists(SELECT 1 FROM information_schema.tables WHERE table_name = $1)"

var queryAddMigration = `insert into migrations (path) values ($1)`

var querySelectWithWheres = "select %s from %s where %s"

var querySelectWithOneWhere = "select %s from %s where %s = $1"

var querySelectWithoutWheres = "select %s from %s"

var queryInsert = "insert into %s (%s) values (%s)"

var queryUpdate = "update %s set %s where %s"
