package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!\n", r.URL.Path[1:])
}

func main() {
	fmt.Println("Server listening on :8080 ...")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
