package gopostgres

//Where objects are used as conditions when building queries
type Where struct {
	column    string
	operation string
	value     interface{}
}

//Returns a "where column = value" condition
func WhereEquals(column string, value interface{}) Where {
	return Where{
		column:    column,
		operation: "=",
		value:     value,
	}
}

//Returns a "where column like %value%" condition
func WhereLike(column string, value string) Where {
	return Where{
		column:    column,
		operation: "like",
		value:     "%"+value+"%",
	}
}
