package gopostgres

import "testing"

func TestWhereLike(t *testing.T) {
	pairsToTest := [][2]string{
		{"col1", "val1"},
		{"col2", "A"},
		{"col3", "23"},
		{"col4", " "},
		{"col5", "NULL"},
		{"col6", "___"},
		{"col7", "سلام"},
	}
	for _, pair := range pairsToTest {
		t.Run(pair[0], func (t *testing.T) {
			col := pair[0]
			val := pair[1]
			wl := WhereLike(col, val)
			if wl.column != col {
				t.Errorf("Expected column of whereLike to be %s, got %s", col, wl.column)
			}
			if wl.operation != "like" {
				t.Errorf("Expected operation of whereLike to be 'like', got %s", wl.operation)
			}
			if wl.value != "%"+val+"%" {
				t.Errorf("Expected value of whereLike to be %s, got %s", val, wl.value)
			}
		})
	}
}

func TestWhereEquals(t *testing.T) {
	pairsToTest := [][2]string{
		{"col1", "val1"},
		{"col2", "A"},
		{"col3", "23"},
		{"col4", " "},
		{"col5", "NULL"},
		{"col6", "___"},
		{"col7", "سلام"},
	}
	for _, pair := range pairsToTest {
		t.Run(pair[0], func (t *testing.T) {
			col := pair[0]
			val := pair[1]
			wl := WhereEquals(col, val)
			if wl.column != col {
				t.Errorf("Expected column of whereEquals to be %s, got %s", col, wl.column)
			}
			if wl.operation != "=" {
				t.Errorf("Expected operation of whereEquals to be 'like', got %s", wl.operation)
			}
			if wl.value != val {
				t.Errorf("Expected value of whereEquals to be %s, got %s", val, wl.value)
			}
		})
	}
}