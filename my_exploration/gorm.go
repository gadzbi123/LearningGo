package main

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type tbl1 struct {
	one string `gorm:"one"`
	two int    `gorm:"two"`
}

var tbl1_values []tbl1

func main() {
	db, err := gorm.Open(sqlite.Open("./test_db.sql"), &gorm.Config{})
	if err != nil {
		fmt.Println("Could not open db", err)
		return
	}
	db.AutoMigrate(&tbl1{})
	// id := db.Create(&tbl1{one: "lol", two: 69})
	// value, _ := db.Get(fmt.Sprint(id))
	// fmt.Println(value)

	res := db.Find(&tbl1_values)
	fmt.Println(res)
	fmt.Println(tbl1_values)
	var my_user tbl1
	db.Table("tbl1").Select("*").Scan(&my_user)
	fmt.Println(my_user.two)
}
