package main

import (
	"fmt"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Println("This is the home page")
	RenderTemplate(w, "home.page.html")
}
func About(w http.ResponseWriter, r *http.Request) {
	fmt.Println("This is the about page")
	RenderTemplate(w, "about.page.html")
}
