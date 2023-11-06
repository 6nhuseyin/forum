package handlers

import (
	"html/template"
	"log"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Home Handler hit")
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	// files := []string {
	// 	"templates/layout.html",
	// 	"templates/home.html",
	// }
	tmpl, err := template.ParseFiles("templates/layout.html", "templates/home.html")
	if err != nil {
		log.Printf("Error Parsing files")
	}
	err = tmpl.ExecuteTemplate(w, "layout", nil)
	if err != nil {
		log.Printf("Error executing template: %v", err)
	}
}
