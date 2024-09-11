package main

import (
	"fmt"
	"strconv"
)

type Product struct {
	Name, Category string
	Price          float64
}

var ProductList = []*Product{
	&Product{"Kajak", "Wodne", 500.50},
	&Product{"Czepek", "Wodne", 50.50},
	&Product{"Klapki", "Wodne", 10.50},
	&Product{"Recznik", "Wodne", 2.5},
	&Product{"Buty", "Górskie", 222.5},
	&Product{"Czapka", "Górskie", 22.5},
	&Product{"Skiety", "Górskie", 5.22},
}

type ProductGroup []*Product

type ProductData map[string]ProductGroup

var Products = make(ProductData)

func ToCurrency(val float64) string {
	return "$" + strconv.FormatFloat(val, 'f', 2, 32)
}
func init() {
	for _, p := range ProductList {
		if _, ok := Products[p.Category]; ok {
			Products[p.Category] = append(Products[p.Category], p)
		} else {
			Products[p.Category] = ProductGroup{p}
		}
	}
}
func (g ProductGroup) TotalPrice(category string, sumReciver chan float64) {
	total := 0.0
	for _, p := range g {
		fmt.Println(category, "product:", p.Name)
		total += p.Price
	}
	sumReciver <- total
}
func CalcStoreTotal(products ProductData) {
	storeTotal := 0.0
	sumReciver := make(chan float64)
	for category, group := range products {
		go func() {
			group.TotalPrice(category, sumReciver)
		}()
	}
	for i := 0; i < len(products); i++ {
		storeTotal += <-sumReciver
	}
	fmt.Println("Total:", ToCurrency(storeTotal))
}
func sender(v float64, ch chan<- float64) {
	ch <- 1.0 + v
}
func receiver(ch <-chan float64) {
	fmt.Println(<-ch)
}
func test() {
	var biDirectionalChan chan float64 = make(chan float64)
	for i := 0; i < 5; i++ {
		go sender(float64(i), biDirectionalChan)
	}
	for i := 0; i < 5; i++ {
		receiver(biDirectionalChan)
	}
}
func addTo2Chans(ch1, ch2 chan *Product) {
	for _, p := range ProductList {
		select {
		case ch1 <- p:
			fmt.Println("Send to ch1", p.Name)
		case ch2 <- p:
			fmt.Println("Send to ch2", p.Name)
		}
	}
	close(ch1)
	close(ch2)
}
func smallChannels() {
	ch1 := make(chan *Product, 2)
	ch2 := make(chan *Product, 2)
	go addTo2Chans(ch1, ch2)

	go func() {
		for v := range ch1 {
			fmt.Println("Received from ch1", v.Name)
		}
	}()
	for v := range ch2 {
		fmt.Println("Received from ch2", v.Name)
	}
}
func main() {
	CalcStoreTotal(Products)
	// test()
	smallChannels()
}
