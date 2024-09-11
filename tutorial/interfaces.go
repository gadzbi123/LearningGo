package main

import (
	"fmt"
)

type Babcia struct {
	wiek      int
	imie      string
	cookSkill int
}

func (v Babcia) String(...any) string {
	return fmt.Sprintf("%v,%v,%v", v.imie, v.wiek, v.cookSkill)
}

type Dziadek struct {
	wiek int
	imie string
}

func (v Dziadek) String(glosno ...any) string {
	// jestGlosno := false
	// if glosno[0].(bool) {
	// 	jestGlosno = true
	// }
	return fmt.Sprintf("%v,%v,%v", v.imie, v.wiek, v.Gadanie(glosno[0].(bool)))
}
func (Dziadek) Gadanie(glosno bool) string {
	if glosno {
		return "DOBRE MURZYN"
	} else {
		return "dobre murzyn"
	}
}

type SpillTheBeans interface {
	String(...any) string
}

func getProductsInfo(products []SpillTheBeans) {
	for _, p := range products {
		switch value := p.(type) {
		case Babcia:
			fmt.Println("Babcia:", value.String())
		case Dziadek:
			fmt.Println("Dziadek:", value.String(true))
		case Point:
			fmt.Println("Point:", value.String())
		}
	}
}

type Point struct {
	x int
	y float32
}

func (p Point) String(...any) string {
	return fmt.Sprintf("{%v,%v}", p.x, p.y)
}

func main() {
	products := []SpillTheBeans{
		Babcia{wiek: 65, imie: "Ola", cookSkill: 4},
		Babcia{55, "Janka", 0},
		Dziadek{55, "Janek"},
		Point{5, 3.3},
	}
	getProductsInfo(products)
}
