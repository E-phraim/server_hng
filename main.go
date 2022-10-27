package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type helloHngResponse struct{
	SlackUsername string `json:"SlackUsername"`
	Backend bool `json:"Backend"`
	Age int `json:"Age"`
	Bio string `json:"Bio"`
}

func helloHng(w http.ResponseWriter, r *http.Request){
	response := helloHngResponse{SlackUsername: "kodeforce98", Backend: true, Age: 24, Bio: "proud firtborn, Golang Developer, committed christian, faithful boyfriend"}
	
	encoder := json.NewEncoder(w)
	encoder.Encode(response)
}

func main(){
	port := 8080

	http.HandleFunc("/hellohng", helloHng)
	log.Printf("Server starting on port %v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}