package main

import (
	"encoding/json"
	"log"
	"net/http"

	"html/template"

	info "github.com/Monishparameswaran/go-microservice/ContainerInfo"
)

type HostInfo struct {
	IP        string `json:"ip"`
	Hostname  string `json:"hostname"`
	Developer string `json:"developer"`
}

func health(w http.ResponseWriter, r *http.Request) {
	hostInfo := HostInfo{
		IP:        info.IP().String(),
		Hostname:  info.Hostname(),
		Developer: "Monish",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(hostInfo)
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		hostInfo := HostInfo{
			IP:        info.IP().String(),
			Hostname:  info.Hostname(),
			Developer: "Monish",
		}

		tmpl, err := template.New("index").Parse(`
<!DOCTYPE html>
<html>
<head>
    <title>Host Information</title>
    <style>
        body {
            font-family: Arial, sans-serif;
            background-color: #f0f0f0;
            display: flex;
            align-items: center;
            justify-content: center;
            height: 100vh;
        }
        .container {
            background-color: #fff;
            border-radius: 5px;
            box-shadow: 0 0 10px rgba(0, 0, 0, 0.2);
            padding: 20px;
            text-align: center;
        }
        h1 {
            color: #0074d9;
        }
        p {
            font-size: 18px;
            margin: 10px 0;
        }
    </style>
</head>
<body>
    <div class="container">
        <h1>Host Information</h1>
        <p>Host IP: {{.IP}}</p>
        <p>Hostname: {{.Hostname}}</p>
        <p>Developed by: {{.Developer}}</p>
    </div>
</body>
</html>
`)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		err = tmpl.Execute(w, hostInfo)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	log.Println("Started serving at http://localhost:8084")
	log.Fatal(http.ListenAndServe(":8084", nil))
}
