package main

import (
	"fmt"
	"net/http"
)

func Printfln(template string, values ...interface{}) {
	fmt.Printf(template+"\n", values...)
}

type Product struct {
	Name, Category string
	Price          float64
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

type StringHandler struct {
	message string
}

func (s StringHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte(s.message))
}

func main() {
	http.Handle("/message", StringHandler{"Hello, world!"})
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.Handle("/*", http.RedirectHandler("/message", http.StatusTemporaryRedirect))
	fsHandler := http.FileServer(http.Dir("./static"))
	http.Handle("/files", http.StripPrefix("/files", fsHandler))
	port := 9090
	fmt.Printf("Listening on port: %v\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%v", port), nil)
	if err != nil {
		fmt.Println("Error accuired", err)
		return
	}

}
