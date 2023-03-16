package main

import (
	"fmt"
	"net/http"
	"text/template"
)

const PORT = ":8080"

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Println("This is the home page")
	renderTemplate(w, "home.page.html")
}
func About(w http.ResponseWriter, r *http.Request) {
	fmt.Println("This is the about page")
	renderTemplate(w, "about.page.html")
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	_ = http.ListenAndServe(PORT, nil)
}

func renderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("Error parsing template ", err)
		return
	}
}
