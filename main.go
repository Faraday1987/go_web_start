package main

import (
	"fmt"
	"net/http"
)

const PORT = ":8080"

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Printf("Server run on port: %s ", PORT)

	_ = http.ListenAndServe(PORT, nil)
}
