package gopostgres

const (
	noRecordFoundMessage = "no record was found"
	moreThanOneRecordFoundMessage = "more than one record was found"
)

//Is returned when no record is found for a specific query
type NoRecordFound struct {
}

func( e NoRecordFound) Error() string {
	return noRecordFoundMessage
}

//Is returned when more than one record is found for a FindBy query
type MoreThanOneRecordFound struct {
}

func( e MoreThanOneRecordFound) Error() string {
	return moreThanOneRecordFoundMessage
}
