package main

import (
	"fmt"
	"os"
	"strings"
	"text/template"
)

type Product struct {
	Name, Category string
	Price          float64
}

var Kayak = Product{
	Name:     "Kayak",
	Category: "Watersports",
	Price:    279,
}

func (p *Product) GetName(upperCase bool) (res string) {
	if upperCase {
		res = strings.ToUpper(p.Name)
	} else {
		res = p.Name
	}
	return
}

var Products = []Product{
	{"Kayak", "Watersports", 279},
	{"Lifejacket", "Watersports", 49.95},
	{"Soccer Ball", "Soccer", 19.50},
	{"Corner Flags", "Soccer", 34.95},
	{"Stadium", "Soccer", 79500},
	{"Thinking Cap", "Chess", 16},
	{"Unsteady Chair", "Chess", 75},
	{"Bling-Bling King", "Chess", 1200},
}

func (p *Product) AddTax() float64 {
	return p.Price * 1.2
}

func (p *Product) ApplyDiscount(amount float64) float64 {
	return p.Price - amount
}

func main() {
	// for _, p := range Products {
	// 	fmt.Printf("Product: %v, Category: %v, Price: $%.2f\n",
	// 		p.Name, p.Category, p.Price)
	// }
	templs, err := template.ParseGlob("templ/*.html")
	if err != nil {
		fmt.Println("Error on parse:", err)
		return
	}
	for _, t := range templs.Templates() {
		err := t.Execute(os.Stdout, &Kayak)
		os.Stdout.WriteString("\n")
		if err != nil {
			fmt.Println(t.Name, "has error:", err)
		}
	}
	t, err := template.ParseFiles("templ/multiple.gohtml")
	if err != nil {
		fmt.Println("Error on multiple:", err)
		return
	}
	t.Execute(os.Stdout, Products)
	t, err = template.ParseFiles("templ/takeSome.gohtml")
	if err != nil {
		fmt.Println("Error on takeSome:", err)
		return
	}
	t.Execute(os.Stdout, Products)
}
