package main

import (
	"fmt"
	"net/http"
	"net/url"
)

type data string

func (d data) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	uri, err := url.Parse(req.RequestURI)
	if err != nil {
		fmt.Fprintln(w, "fucked", err)
		return
	}
	q, err := url.ParseQuery(uri.RawQuery)
	if err != nil {
		fmt.Fprintln(w, "fucked 2", err)
		return
	}
	val := q["key"]
	cookie, err := req.Cookie("key")
	if len(val) == 0 && err != nil {
		fmt.Fprintln(w, "fucked 3", err)
		return
	}
	if cookie != nil {
		fmt.Fprint(w, " data ", d, " from ", cookie)
		return
	}
	c := &http.Cookie{Name: "key", Value: val[0], MaxAge: 60}
	http.SetCookie(w, c)
	fmt.Fprint(w, " data ", d, " from ", req.Cookies())
}
func main() {
	fmt.Println("Serving on port 8080")
	mux := http.NewServeMux()
	mux.Handle("/dupą", data("XD"))
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Println("Server error:", err)
	}
}
