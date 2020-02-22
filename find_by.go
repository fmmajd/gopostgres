package gopostgres

import (
	"fmt"
)

// Searches for a single row with specific column value, and returns the row
//
// If no row is found, a NoRecordFound error will be returned
//
// If more than one row is found, a MoreThanOneRecordFound error will be returned
func (db Postgres) FindBy(tableName string, column string, value interface{}) ([]interface{}, error){
	q := fmt.Sprintf(querySelectWithOneWhere, "*", tableName, column)
	params := []interface{}{
		value,
	}

	query := query{
		Statement: q,
		Args:      params,
	}
	res, err := db.execQuery(query)
	if err != nil {
		return nil, err
	}

	if len(res) > 1 {
		return nil, MoreThanOneRecordFound{}
	}
	if len(res) == 0 {
		return nil, NoRecordFound{}
	}

	return res[0], nil
}
