package main

import (
	"log"
	"net/http"
	"fmt"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/index.html")
	})
	log.Println("Started serving at http://localhost:8082")
	http.ListenAndServe(":8082", nil)
	log.Fatal(http.ListenAndServe(":8082", nil))
}
