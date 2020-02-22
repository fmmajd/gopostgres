package gopostgres

func (db Postgres) tableExists(name string) (bool, error) {
	q := query{
		Statement: queryTableExists,
		Args:      []interface{}{name},
	}
	res, err := db.execQuery(q)
	if err != nil {
		return false, err
	}

	exists := res[0][0]

	return exists.(bool), nil
}
