package main

import (
	"database/sql"
	"fmt"
	"reflect"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./test_db.sql")
	if err != nil {
		fmt.Println("Error on opening db", err)
		return
	}
	var i int
	err = db.QueryRow("insert into tbl1 values('abc',10);").Scan(&i)
	switch err {
	case sql.ErrNoRows:
	default:
		fmt.Println("Error on insert query", err)
		return
	}
	fmt.Println("Added successfully")
	rows, err := db.Query("select * from tbl1;")
	if err != nil {
		fmt.Println("Error on tbl1 read", err)
		return
	}
	db_types, err := rows.ColumnTypes()
	if err != nil {
		fmt.Println("Failed to get types", err)
		return
	}
	for i, t := range db_types {
		Type := t.ScanType()
		fmt.Println(i, Type, reflect.TypeOf(Type).Kind())
	}
	for rows.Next() {
	}
	err = db.Close()
	if err != nil {
		fmt.Println("Error on close db", err)
		return
	}

}
