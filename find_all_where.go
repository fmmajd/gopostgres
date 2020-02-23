package gopostgres

import (
	"fmt"
	"strings"
)

//Use this function to find rows by multiple where conditions
func (db Postgres) FindAllWhere(tableName string, columns []string, conditions ...Where) ([][]interface{}, error){
	var q string
	selects := strings.Join(columns, ",")
	var params []interface{}
	if len(conditions) > 0 {
		whereStatement := ""
		for i,w := range conditions {
			if i!= 0 {
				whereStatement += " and"
			}
			whereStatement += fmt.Sprintf("%s %s $%d", w.column, w.operation, i+1)
			params = append(params, w.value)
		}
		q = fmt.Sprintf(querySelectWithWheres, selects, tableName, whereStatement)
	} else {
		q = fmt.Sprintf(querySelectWithoutWheres, selects, tableName)
	}

	query := query{
		Statement: q,
		Args:      params,
	}
	res, err := db.execQuery(query)
	if err != nil {
		return nil, err
	}

	return res, nil
}
