package main

import (
	"fmt"
	"reflect"
	"strings"
)

type Customer struct {
	Name, City string
}
type Product struct {
	Name, Category string
	Price          float64
}
type Payment struct {
	Currency string
	Amount   float64
}

func PrintCustomersAndProducts(cps ...interface{}) {
	for _, cp := range cps {
		switch val := cp.(type) {
		case Product:
			fmt.Printf("Product: Name - %v, Category - %v, Price: %.2f\n", val.Name, val.Category, val.Price)
		case Customer:
			fmt.Printf("Product: Name - %v, City - %v\n", val.Name, val.City)
		}
	}
}

func printDetails(values ...any) {
	for _, elem := range values {
		details := []string{}
		elemType := reflect.TypeOf(elem)
		elemValue := reflect.ValueOf(elem)
		switch elemType.Kind() {
		case reflect.Struct:
			for i := 0; i < elemType.NumField(); i++ {
				elName := elemType.Field(i).Name
				elVal := elemValue.Field(i)
				details = append(details, fmt.Sprintf("%v: %v", elName, elVal))
			}
			fmt.Printf("%v := %v\n", elemType.Name(), strings.Join(details, ", "))
		default:
			fmt.Printf("%v: %v\n", elemType.Name(), elemValue)
		}
	}
}
func main() {
	p1 := Product{Name: "Ball", Category: "Watersports", Price: 22.6}
	c1 := Customer{Name: "Antek", City: "NY"}
	pm1 := Payment{Currency: "Zł", Amount: 25.5}
	PrintCustomersAndProducts(p1, c1)
	printDetails(p1, c1, true, 10, pm1)

}
