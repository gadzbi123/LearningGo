package main

import (
	"database/sql"
	"fmt"

	// _ "ncruces/go-sqlite3"

	_ "github.com/ncruces/go-sqlite3/driver"
	_ "github.com/ncruces/go-sqlite3/embed"
)

type Category struct {
	id   int
	name string
}
type Product struct {
	id int
	Category
	name  string
	price sql.NullFloat64
}

func (p Product) String() string {
	if p.price.Valid {
		return fmt.Sprintf("%v - %v - %v - $%.2f", p.id, p.name, p.Category.name, p.price.Float64)
	}
	return fmt.Sprintf("%v - %v - %v - No price", p.id, p.name, p.Category.name)

}

func getAllCategorized(db *sql.DB) {
	rows, err := db.Query(`SELECT p.Id, p.Name, p.Price, c.Id, c.Name 
	FROM Products p, Categories c 
	WHERE p.Category = c.Id`)
	if err != nil {
		fmt.Println("Field to do query", err)
		return
	}

	products := []Product{}
	for rows.Next() {
		p := Product{}
		err := rows.Scan(&p.id, &p.name, &p.price, &p.Category.id, &p.Category.name)
		if err != nil {
			fmt.Println("Error on db scan", err)
			continue
		}
		products = append(products, p)
	}
	fmt.Println("--------------------------ALL PRODUCTS--------------------------")
	fmt.Println(products)
}

func getSpecificCategory(db *sql.DB, categoryName string) {

	rows, err := db.Query(`SELECT p.Id, p.Name, p.Price, c.Id, c.Name 
	FROM Products p, Categories c
	WHERE c.Id = p.Category AND c.Name = ?
	`, categoryName)
	if err != nil {
		fmt.Println("Failed to query db", err)
	}
	for rows.Next() {
		p := Product{}
		err := rows.Scan(&p.id, &p.name, &p.price, &p.Category.id, &p.Category.name)
		if err != nil {
			fmt.Println("Failed to scan row", err)
		}
		fmt.Printf("Product: %v\n", p)
	}
}

func transactionProduct(db *sql.DB, p Product) error {
	tx, _ := db.Begin()
	query, err := db.Prepare("INSERT INTO Products (Name, Price, Category) VALUES (?, ?, ?)")
	if err != nil {
		tx.Rollback()
		return err
	}
	res, err := tx.Stmt(query).Exec(p.name, p.price, p.Category.id)
	if err != nil {
		tx.Rollback()
		return err
	}
	id, err := res.LastInsertId()
	if err != nil {
		tx.Rollback()
		return err
	}
	fmt.Println("id of new product", id)
	tx.Commit()
	fmt.Println("transaction successful")
	return nil

}

// Hydrate db
// sqlite3 "products.db" ".read products.sql"
func main() {
	db, err := sql.Open("sqlite3", "products.db")
	if err != nil {
		fmt.Println("Failed to open db", err)
		return
	}
	defer db.Close()
	getSpecificCategory(db, "Soccer")
	p := Product{name: "Maska", price: sql.NullFloat64{}, Category: Category{id: 2}}
	err = transactionProduct(db, p)
	if err != nil {
		fmt.Println("Failed to make transaction", err)
	}
	// getAllCategorized(db)
	getSpecificCategory(db, "Soccer")
}
