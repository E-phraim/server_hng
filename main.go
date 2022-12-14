package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

type helloHngResponse struct {
	SlackUsername string `json:"slackUsername"`
	Backend       bool   `json:"backend"`
	Age           int    `json:"age"`
	Bio           string `json:"bio"`
}

func helloHng(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	w.Header().Add("Content-Type", "application/json")
	response := helloHngResponse{SlackUsername: "kodeforce98", Backend: true, Age: 24, Bio: "Proud firstborn, Golang Developer, Committed christian, Faithful boyfriend"}

	encoder := json.NewEncoder(w)
	encoder.Encode(response)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "9090"
	}

	http.HandleFunc("/", helloHng)
	log.Printf("Server starting on port %v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}
