package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8008", nil)
}

func handler(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "Hello World!")
}
