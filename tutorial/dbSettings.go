package main

import (
	"database/sql"
	"fmt"
	"time"
)

func main() {
	db, err := sql.Open("mysql", "test.db")
	if err != nil {
		fmt.Println("FAIL HERE", err)
	}
	db.SetMaxIdleConns(10) // idle contains in open conns
	db.SetMaxOpenConns(20)
	db.SetConnMaxIdleTime(time.Minute)      // remove idle conns after time
	db.SetConnMaxLifetime(10 * time.Minute) // remove connection after time
}
