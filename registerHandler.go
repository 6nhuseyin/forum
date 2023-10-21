package main

import (
	"log"
	"net/http"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		log.Print("Wrong Method")
		http.Error(w, "wrong method", http.StatusBadRequest)
	}

	err := r.ParseForm()
	if err != nil{
		log.Print("Couldn't Parse Form")
		http.Error(w, "server couldn't parse the form", http.StatusInternalServerError)
	}


	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	email := r.PostFormValue("email")

	

}