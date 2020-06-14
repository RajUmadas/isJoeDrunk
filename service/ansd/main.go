package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4"
)

type ans struct {
	Ans       string `json:"ans"`
	Lit_Level string `json:"lit_level"`
}

func getAnsDb() (string, string, error) {
	chanelEnv := os.Getenv("DEPLOY_CHANNEL")
	if chanelEnv == "local-devel" {
		conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
			os.Exit(1)
		}
		defer conn.Close(context.Background())

		var db_ans string
		var db_lit_level string
		err = conn.QueryRow(context.Background(), "select ans, lit_level from ans order by date_created DESC LIMIT 1;").Scan(&db_ans, &db_lit_level)

		if err != nil {
			return "0", "0", err
		}

		return db_ans, db_lit_level, nil

	}
	return "0", "0", nil
}

func getAns(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ansd get handler")
	x, y, err := getAnsDb()
	if err != nil {
		os.Exit(1)
	}
	ans := ans{
		Ans:       x,
		Lit_Level: y,
	}
	json.NewEncoder(w).Encode(ans)
}

func main() {

	r := mux.NewRouter()
	// Routes consist of a path and a handler function.
	r.HandleFunc("/", getAns)
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	// Bind to a port and pass our router in
	//log.Fatal(http.ListenAndServe(":8000", r))
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(originsOk, headersOk, methodsOk)(r)))
}
