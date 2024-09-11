package main

import (
	"bytes"
	"fmt"
	"io"
	"strings"
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

func Printfln(template string, values ...interface{}) {
	fmt.Printf(template+"\n", values...)
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

func processData(reader io.Reader, writer io.Writer) {
	buffer := make([]byte, 4)
	for {
		n, err := reader.Read(buffer)
		if n > 0 {
			Printfln("Bytes read: \"%v\", bytes received: %v, err: %v", string(buffer[:n]), n, err)
			writer.Write(buffer[:n])
		}
		if err == io.EOF {
			break
		}
	}
}
func main() {
	Printfln("Kayak name: %v, Kayak price: %v", Kayak.Name, Kayak.Price)
	reader := strings.NewReader("jedzenie1")
	var writer bytes.Buffer
	processData(reader, &writer)
	Printfln("Writer: %v", writer.String())
}
