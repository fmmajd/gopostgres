package gopostgres

import (
	"fmt"
	"strconv"
)

//This function saves a new record in the database.
//IMPORTANT: new records MUST return 0 for function PostgresId(), otherwise a  NewRecordWithUnZeroId error would return
func (db Postgres) Insert(r Record) (int64, error) {
	if r.PostgresId() != 0 {
		return 0, NewRecordWithUnZeroId{}
	}

	columns := ""
	values := ""
	var params []interface{}
	i := 0
	for k, v := range r.PostgresValues() {
		if i != 0 {
			columns += ","
			values += ","
		}
		columns += k
		values += "$"+strconv.Itoa(i+1)
		params = append(params, v)
		i++
	}

	q := fmt.Sprintf(queryInsert, r.PostgresTable(), columns, values)

	qu := query{
		Statement: q,
		Args:      params,
	}

	res, err := db.execQuery(qu)
	if err != nil {
		return 0, err
	}
	id := res[0][0]

	return id.(int64), nil
}
