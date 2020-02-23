package gopostgres

import (
	"github.com/jackc/pgx/v4"
	"reflect"
	"testing"
	"time"
)

func TestFindBy(t *testing.T) {
	InitDB("test", "user", "pass", "localhost", nil, pgx.LogLevelDebug)
	tableExists, _ := DB.tableExists(testTable)
	if tableExists {
		purgeTableQuery := query{
			Statement: truncateTestTable,
		}
		DB.execQuery(purgeTableQuery)
	} else {
		createTestTableQuery := query{
			Statement: createTestTable,
		}
		DB.execQuery(createTestTableQuery)
	}
	testUsername := "username1"
	testUpdateTime := time.Date(2000, 1, 1, 13, 10, 20, 0, time.UTC)
	//No record should be found first
	res, err := DB.FindBy(testTable, "username", testUsername)
	if err == nil {
		t.Error("Expected an error, got nothing")
	}else {
		switch err.(type) {
		case NoRecordFound:
			//
		default:
			t.Errorf("Expected to get a NoRecordFound error, got %s", reflect.TypeOf(err))
		}
	}
	if res != nil {
		t.Errorf("Expected to receive nil result, got %v", res)
	}

	//now we insert a single record
	insertQuery := query{
		Statement: insertOneRowInTestTable,
		Args: []interface{} {
			testUsername,
			testUpdateTime,
		},
	}
	DB.execQuery(insertQuery)

	//the single record should be found
	res, err = DB.FindBy(testTable, "username", testUsername)
	if err != nil {
		t.Errorf("Expected to get no error, got %s", err.Error())
	}
	receivedUsername := res[1].(string)
	if receivedUsername != testUsername {
		t.Errorf("Expected the found username to be %s, got %s", testUsername, receivedUsername)
	}
	receivedUpdateTime := res[2].(time.Time)
	if receivedUpdateTime != testUpdateTime {
		t.Errorf("Expected the found update time to be %v, got %v", testUpdateTime, receivedUpdateTime)
	}

	insertQuery.Args = []interface{}{"username2", testUpdateTime}
	DB.execQuery(insertQuery)
	res, err = DB.FindBy(testTable, "last_updated", testUpdateTime)
	if err == nil {
		t.Error("Expected an error, got nothing")
	}else {
		switch err.(type) {
		case MoreThanOneRecordFound:
			//
		default:
			t.Errorf("Expected to get a MoreThanOneRecordFound error, got %s", reflect.TypeOf(err))
		}
	}
	if res != nil {
		t.Errorf("Expected to receive nil result, got %v", res)
	}
}
