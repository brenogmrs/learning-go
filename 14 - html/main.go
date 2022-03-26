package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var templates *template.Template

type user struct {
	Name  string
	Email string
}

func main() {

	u := user{
		"Jhon",
		"jhondoe@email.com",
	}

	templates = template.Must(template.ParseGlob("*.html"))

	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		templates.ExecuteTemplate(w, "index.html", u)
	})

	fmt.Println("Server running on port 5000")
	log.Fatal(http.ListenAndServe(":5000", nil))
}
