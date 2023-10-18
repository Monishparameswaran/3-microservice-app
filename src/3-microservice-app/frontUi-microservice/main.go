package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		// Redirect to the Hello microservice running on localhost:1002
		http.Redirect(w, r, "helloui:8082", http.StatusSeeOther) 
	})

	http.HandleFunc("/hostinfo", func(w http.ResponseWriter, r *http.Request) {
		// Redirect to the HostInfo microservice running on localhost:1003
		http.Redirect(w, r, "hostinfo:8084", http.StatusSeeOther)
	})
	log.Println("Started serving at http://localhost:8080")
	log.Println("Please make sure all containers are in custom bridge network in case of Docker");
	http.ListenAndServe(":8080", nil)
}
