package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/index.html")
	})
	http.ListenAndServe(":8082", nil)
	log.Fatal(http.ListenAndServe(":8082", nil))
}
