package gopostgres

import (
	"fmt"
)

//Returns all the rows by given condition
//
//A NoRecordFound error will be returned if no row was matched
func (db Postgres) FindAllBy(tableName string, column string, value interface{}) ([][]interface{}, error){
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

	if len(res) == 0 {
		return nil, NoRecordFound{}
	}

	return res, nil
}
