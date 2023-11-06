package main

import (
	"fmt"
	"log"
	"net/http"

	"forum/data-access"
	"forum/router"
)

// var tmpl *template.Template

func init() {
	err := data-access.InitDatabase()
	
	if err != nil {
		log.Fatalf("Failed to connect to database: %v\n", err)
	}
	fmt.Println("Success connected to database")

	
}



func Server() {

	fs := http.FileServer(http.Dir("static"))

	//initialize a new servemux and register the home function as the handler for the "/" URL pattern.
	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static/", fs))
	mux.HandleFunc("/", HomeHandler)

	//Start a new web server listening on port 7000
	log.Print("Starting server on :7001")
	err := http.ListenAndServe(":7001", mux)
	if err != nil {
		log.Fatal(err)
	}

}

func pathHandler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		HomeHandler(w, r)
	case "/register":
		RegisterHandler(w, r)
	case "/login":
		LoginHandler(w, r)
	default:
		ErrorHandler(w, r, 404)
	}

}

func main() {
	Server()
	http.HandleFunc("/", pathHandler)
}
