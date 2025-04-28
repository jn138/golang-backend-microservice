package mysql_test

import (
	"fmt"
	"golang-backend-microservice/dataservice/mysql"
	"golang-backend-microservice/model"
	"reflect"
	"testing"
)

func TestBuildQueryColumns(t *testing.T) {
	var (
		expectedType  = "[]mysql.column"
		expectedName  = "foo"
		expectedValue = "bar"
	)
	params := map[string]interface{}{expectedName: expectedValue}
	columns := mysql.BuildQueryColumns(params)

	if fmt.Sprint(reflect.TypeOf(columns)) != expectedType {
		t.Errorf("Expected %v, got %v: incorrect type", expectedType, fmt.Sprint(reflect.TypeOf(columns)))
	}
	if columns[0].Name != expectedName {
		t.Errorf("Expected %v, got %v: incorrect name", expectedName, columns[0].Name)
	}
	if columns[0].Value != expectedValue {
		t.Errorf("Expected %v, got %v: incorrect value", expectedValue, columns[0].Value)
	}
}

func TestBuildSelectQuery(t *testing.T) {
	var (
		expectedQ     = "SELECT * FROM books WHERE title = ? LIMIT 1"
		expectedTitle = "Harry Potter"
	)
	args := model.MySqlReqArgs{
		Table: "books",
		Where: map[string]interface{}{
			"title": expectedTitle,
		},
		Limit: "1",
	}
	q, d, err := mysql.BuildSelectQuery(&args)

	if err != nil {
		t.Errorf("Error not nil: %v", err)
	}
	if q != expectedQ {
		t.Errorf("Expected %v, got %v: incorrect query", expectedQ, q)
	}
	if d[0] != expectedTitle {
		t.Errorf("Expected %v, got %v: incorrect argument", expectedTitle, d[0])
	}

	var (
		expectedQ2     = "SELECT * FROM books WHERE genre IN (?) AND year IN (?,?,?)"
		expectedGenres = []string{"fiction"}
		expectedYears  = []string{"2020", "2021", "2022"}
	)
	args = model.MySqlReqArgs{
		Table: "books",
		Where: map[string]interface{}{
			"genre": expectedGenres,
			"year":  expectedYears,
		},
	}
	q, d, err = mysql.BuildSelectQuery(&args)

	if err != nil {
		t.Errorf("Error not nil: %v", err)
	}
	if q != expectedQ2 {
		t.Errorf("Expected %v, got %v: incorrect query", expectedQ2, q)
	}
	if d[0] != expectedGenres[0] {
		t.Errorf("Expected %v, got %v: incorrect argument", expectedGenres[0], d[0])
	}
	if d[1] != expectedYears[0] {
		t.Errorf("Expected %v, got %v: incorrect argument", expectedYears[0], d[1])
	}
}

func TestBuildInsertQuery(t *testing.T) {
	var (
		expectedQ     = "INSERT INTO books (title) VALUES (?)"
		expectedTitle = "The Replublic"
	)
	args := model.MySqlReqArgs{
		Table: "books",
		Data: map[string]interface{}{
			"title": expectedTitle,
		},
	}
	q, d, err := mysql.BuildInsertQuery(&args)

	if err != nil {
		t.Errorf("Error not nil: %v", err)
	}
	if q != expectedQ {
		t.Errorf("Expected %v, got %v: incorrect query", expectedQ, q)
	}
	if d[0] != expectedTitle {
		t.Errorf("Expected %v, got %v: incorrect argument", expectedTitle, d[0])
	}
}

func TestBuildUpdateQuery(t *testing.T) {
	var (
		expectedQ     = "UPDATE books SET year = ? WHERE title = ?"
		expectedYear  = "-234"
		expectedTitle = "The Republic"
	)
	args := model.MySqlReqArgs{
		Table: "books",
		Where: map[string]interface{}{"title": expectedTitle},
		Data:  map[string]interface{}{"year": expectedYear},
	}
	q, d, err := mysql.BuildUpdateQuery(&args)

	if err != nil {
		t.Errorf("Error not nil: %v", err)
	}
	if q != expectedQ {
		t.Errorf("Expected %v, got %v: incorrect query", expectedQ, q)
	}
	if d[0] != expectedYear {
		t.Errorf("Expected %v, got %v: incorrect argument", expectedYear, d[0])
	}
	if d[1] != expectedTitle {
		t.Errorf("Expected %v, got %v: incorrect argument", expectedTitle, d[1])
	}

	var (
		expectedQ2     = "UPDATE books SET author = ? WHERE title IN (?)"
		expectedAuthor = "JK Rowling"
		expectedTitles = []string{"Harry Potter"}
	)
	args = model.MySqlReqArgs{
		Table: "books",
		Where: map[string]interface{}{"title": expectedTitles},
		Data:  map[string]interface{}{"author": expectedAuthor},
	}
	q, d, err = mysql.BuildUpdateQuery(&args)

	if err != nil {
		t.Errorf("Error not nil: %v", err)
	}
	if q != expectedQ2 {
		t.Errorf("Expected %v, got %v: incorrect query", expectedQ2, q)
	}
	if d[0] != expectedAuthor {
		t.Errorf("Expected %v, got %v: incorrect argument", expectedAuthor, d[0])
	}
	if d[1] != expectedTitles[0] {
		t.Errorf("Expected %v, got %v: incorrect argument", expectedTitles[0], d[1])
	}
}

func TestBuildDeleteQuery(t *testing.T) {
	var (
		expectedQ      = "DELETE FROM books WHERE author = ?"
		expectedAuthor = "JK Rowling"
	)
	args := model.MySqlReqArgs{
		Table: "books",
		Where: map[string]interface{}{"author": expectedAuthor},
	}
	q, d, err := mysql.BuildDeleteQuery(&args)

	if err != nil {
		t.Errorf("Error not nil: %v", err)
	}
	if q != expectedQ {
		t.Errorf("Expected %v, got %v: incorrect query", expectedQ, q)
	}
	if d[0] != expectedAuthor {
		t.Errorf("Expected %v, got %v: incorrect argument", expectedAuthor, d[0])
	}
}
