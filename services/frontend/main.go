package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"text/template"
)

var backend = ""

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Handling %+v\n", r)
	resp, herr := http.Get(backend)

	if herr != nil {
		http.Error(w, fmt.Sprintf("Error communicating with backend: %v", herr), 500)
		return
	}

	if resp.StatusCode != http.StatusOK {
		http.Error(w, fmt.Sprintf("Backend seems unhealthy: %v", resp), 500)
		return
	}

	body, ierr := ioutil.ReadAll(resp.Body)

	if ierr != nil {
		http.Error(w, fmt.Sprintf("Backend producing garbage: %v", ierr), 500)
		return
	}

	host, err := os.Hostname()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error retrieving hostname: %v", err), 500)
		return
	}
	msg2 := fmt.Sprintf("%s", host)

	b, err := ioutil.ReadFile("/opt/config/data/config.properties")
	if err != nil {
		http.Error(w, fmt.Sprintf("Error reading configuration file hostname: %v", err), 500)
	}
	configuration := string(b)

	t, terr := template.ParseFiles("/content/index.html")

	if terr != nil {
		http.Error(w, fmt.Sprintf("Error loading template: %v", terr), 500)
		return
	}

	config := map[string]string{
		"Message":      string(body),
		"Feature":      os.Getenv("FEATURE"),
		"FrontendHost": string(msg2),
		"Config":       Configuration,
	}

	t.Execute(w, config)
}

func main() {
	backend = os.Getenv("BACKEND_ENDPOINT")
	if backend == "" {
		fmt.Printf("Backend not supplied, use BACKEND_ENDPOINT in env")
		os.Exit(1)
	}

	http.HandleFunc("/", index)
	http.HandleFunc("/backend", index)

	port := ":8000"
	fmt.Printf("Starting to service on port %s\n", port)
	http.ListenAndServe(port, nil)
}
