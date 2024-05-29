package main

import (
	"fmt"
	"net/http"
)

type counter int

func (ctr *counter) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	*ctr++
	if *ctr%1000 == 0 {
		fmt.Printf("Called counter: %v\n", *ctr)
	}
	fmt.Fprintf(w, "counter=%v\n", *ctr)
}

type Chan chan *http.Request

func (c Chan) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c <- r
	fmt.Println("Notification sent:", w)
}
func main() {
	ctr := new(counter)
	// var w http.ResponseWriter
	mux := http.NewServeMux()
	mux.Handle("/ctr", ctr)
	http.ListenAndServe(":8080", mux)
	// http.ListenAndServe("127.0.0.1:8080", ctr)

}
