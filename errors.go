package gopostgres

const (
	noRecordFoundMessage = "no record was found"
	moreThanOneRecordFoundMessage = "more than one record was found"
	NewRecordWithUnZeroIdMessage = "new records must have an id of zero"
	UpdatingRecordWithZeroIdMessage = "updating records must have a non-zero id"
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

type NewRecordWithUnZeroId struct {
}

func (e NewRecordWithUnZeroId) Error() string {
	return NewRecordWithUnZeroIdMessage
}

type UpdatingRecordWithZeroId struct {
}

func (e UpdatingRecordWithZeroId) Error() string {
	return UpdatingRecordWithZeroIdMessage
}
