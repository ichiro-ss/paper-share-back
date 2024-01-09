package main

import (
	"fmt"
	"net/http"
)

func main() {
	h1 := func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprintln(w, "Hello World!")
	}

	http.HandleFunc("/", h1)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
