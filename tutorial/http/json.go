package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"text/template"
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

type Context struct {
	req  *http.Request
	Data []Product
}

var ht *template.Template

func HandleTemplateRequest(w http.ResponseWriter, req *http.Request) {
	path := req.URL.Path
	if path == "" {
		path = "products.html"
	}
	t := ht.Lookup(path)
	if t == nil {
		http.NotFoundHandler()
	} else {
		err := ht.Execute(w, Context{req: req, Data: Products})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func HandleJsonRequest(w http.ResponseWriter, req *http.Request) {
	req.Header.Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Products)
}

func main() {
	var err error
	ht.Funcs(map[string]interface{}{
		"intVal": strconv.Atoi,
	})
	ht, err = template.ParseGlob("templates/*.html")
	if err != nil {
		fmt.Println("Error on parsing:", err)
		return
	}
	http.Handle("/templates/", http.StripPrefix("/templates/", http.HandlerFunc(HandleTemplateRequest)))
	http.HandleFunc("/json", HandleJsonRequest)
	port := 9090
	fmt.Printf("Listening on port: %v\n", port)
	err = http.ListenAndServe(fmt.Sprintf(":%v", port), nil)
	if err != nil {
		fmt.Println("Error accuired", err)
		return
	}

}
