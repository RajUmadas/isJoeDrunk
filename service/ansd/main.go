package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type ans struct {
	Ans   string `json:"ans"`
	Level int    `json:"level"`
}

func getAns(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ansd get handler")
	ans := ans{
		Ans:   "Yes",
		Level: 10,
	}
	json.NewEncoder(w).Encode(ans)
}

func main() {

	r := mux.NewRouter()
	// Routes consist of a path and a handler function.
	r.HandleFunc("/", getAns)

	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8000", r))
}
