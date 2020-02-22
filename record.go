package gopostgres

//implement this interface on types that correspond a table in the postgres database
type Record interface {
	PostgresTable() string
	PostgresColumns() []string
	PostgresValue(column string) interface{}
	PostgresValues() map[string]interface{}
	PostgresId() uint
}
