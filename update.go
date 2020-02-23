package gopostgres

import (
	"fmt"
	"strconv"
)

//Updates a record in the database.
//The primary key is returned with the PostgresId() function and MUST be non-zero, otherwise a UpdatingRecordWithZeroId error is returned
func (db Postgres) Update(r Record) error {
	if r.PostgresId() == 0 {
		return UpdatingRecordWithZeroId{}
	}

	setStatement := ""
	var params []interface{}
	i := 0
	for k, v := range r.PostgresValues() {
		if i!= 0 {
			setStatement += ","
		}
		setStatement += k+"=$"+strconv.Itoa(i+1)
		params = append(params, v)
		i++
	}

	whereStatement := "id=$"+strconv.Itoa(i+1)
	params = append(params, r.PostgresId())

	statement := fmt.Sprintf(queryUpdate, r.PostgresTable(), setStatement, whereStatement)

	qu := query{
		Statement: statement,
		Args:      params,
	}

	_, err := db.execQuery(qu)
	if err != nil {
		return err
	}

	return nil
}
