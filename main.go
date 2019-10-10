package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	var name string
	found := r.URL.Query().Get("name")
	if found != "" {
		name = found
	} else {
		name = "Jenkins"
	}
	fmt.Fprintf(w, "Welcome to  %s!", name)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":80", nil)
}
