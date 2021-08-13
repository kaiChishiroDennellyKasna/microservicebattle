package main

import (
	"encoding/json"
	"fmt"
	"log"
	rand2 "math/rand"
	"net/http"
	"os"
)

func main() {
	port := "8080"
	if v := os.Getenv("PORT"); v != "" {
		port = v
	}
	http.HandleFunc("/", handler)
	http.HandleFunc("/healthcheck", healthcheck)

	log.Printf("starting server on port :%s", port)
	err := http.ListenAndServe(":"+port, nil)
	log.Fatalf("http listen error: %v", err)
}


func healthcheck(w http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	fmt.Fprint(w, "Gucci")
}

func handler(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprint(w, "only POST method supported")
		return
	}
	var v ArenaUpdate
	defer req.Body.Close()
	d := json.NewDecoder(req.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&v); err != nil {
		log.Printf("WARN: failed to decode ArenaUpdate in response body: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	resp := play(v)
	fmt.Fprint(w, resp)
}

func play(input ArenaUpdate) (response string) {
	log.Printf("IN: %#v", input)

	commands := []string{"F", "R", "L", "T"}
	rand := rand2.Intn(4)
	return commands[rand]
}
