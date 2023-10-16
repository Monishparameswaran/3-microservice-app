package main

import (
    "net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, "static/index.html")
    })

    http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
        // Redirect to the Hello microservice running on localhost:1002
        http.Redirect(w, r, "http://localhost:8082", http.StatusSeeOther)
    })

    http.HandleFunc("/hostinfo", func(w http.ResponseWriter, r *http.Request) {
        // Redirect to the HostInfo microservice running on localhost:1003
        http.Redirect(w, r, "http://localhost:8084", http.StatusSeeOther)
    })

    http.ListenAndServe(":8080", nil)
}