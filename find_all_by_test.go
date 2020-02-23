package gopostgres

import (
	"fmt"
	"github.com/jackc/pgx/v4"
	"reflect"
	"testing"
	"time"
)

func TestFindAllBy(t *testing.T) {
	InitDB("test", "user", "pass", "localhost", nil, pgx.LogLevelDebug)
	tableExists, _ := DB.tableExists(testTableForFindAllBy)
	if tableExists {
		purgeTableQuery := query{
			Statement: fmt.Sprintf(truncateTestTable, testTableForFindAllBy),
		}
		DB.execQuery(purgeTableQuery)
	} else {
		createTestTableQuery := query{
			Statement: fmt.Sprintf(createTestTable, testTableForFindAllBy),
		}
		DB.execQuery(createTestTableQuery)
	}
	testUsername := "username1"
	testUpdateTime := time.Date(2000, 1, 1, 13, 10, 20, 0, time.UTC)
	//No record should be found first
	res, err := DB.FindAllBy(testTableForFindAllBy, "username", testUsername)
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
		Statement: fmt.Sprintf(insertOneRowInTestTable, testTableForFindAllBy),
		Args: []interface{} {
			testUsername,
			testUpdateTime,
		},
	}
	DB.execQuery(insertQuery)

	//the single record should be found
	res, err = DB.FindAllBy(testTableForFindAllBy, "username", testUsername)
	if err != nil {
		t.Errorf("Expected to get no error, got %s", err.Error())
	}
	if len(res) != 1 {
		t.Errorf("Expected the found result to be of length %d, got %d", 1, len(res))
	}
	receivedUsername := res[0][1].(string)
	if receivedUsername != testUsername {
		t.Errorf("Expected the found username to be %s, got %s", testUsername, receivedUsername)
	}
	receivedUpdateTime := res[0][2].(time.Time)
	if receivedUpdateTime != testUpdateTime {
		t.Errorf("Expected the found update time to be %v, got %v", testUpdateTime, receivedUpdateTime)
	}

	testUsername2 := "username2"
	insertQuery.Args = []interface{}{testUsername2, testUpdateTime}
	DB.execQuery(insertQuery)
	res, err = DB.FindAllBy(testTableForFindAllBy, "last_updated", testUpdateTime)
	if err != nil {
		t.Errorf("Expected no error, got %s", err.Error())
	}
	if len(res) != 2 {
		t.Errorf("Expected the found result to be of length %d, got %d", 2, len(res))
	}
	receivedUsername1 := res[0][1].(string)
	if receivedUsername1 != testUsername {
		t.Errorf("Expected the found username to be %s, got %s", testUsername, receivedUsername)
	}
	receivedUpdateTime1 := res[0][2].(time.Time)
	if receivedUpdateTime1 != testUpdateTime {
		t.Errorf("Expected the found update time to be %v, got %v", testUpdateTime, receivedUpdateTime)
	}
	receivedUsername2 := res[1][1].(string)
	if receivedUsername2 != testUsername2 {
		t.Errorf("Expected the found username to be %s, got %s", testUsername2, receivedUsername2)
	}
	receivedUpdateTime2 := res[1][2].(time.Time)
	if receivedUpdateTime2 != testUpdateTime {
		t.Errorf("Expected the found update time to be %v, got %v", testUpdateTime, receivedUpdateTime2)
	}
}
