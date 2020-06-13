package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4"
)

func main() {
	chanelEnv := os.Getenv("DEPLOY_CHANNEL")
	if chanelEnv == "local-devel" {
		conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
			os.Exit(1)
		}
		defer conn.Close(context.Background())

		var ans string
		var lit_level string
		err = conn.QueryRow(context.Background(), "select ans, lit_level from ans").Scan(&ans, &lit_level)
		if err != nil {
			fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
			os.Exit(1)
		}
		fmt.Println(ans)
		fmt.Println(lit_level)

		return
	}

	fmt.Println("Dont know the env yo")

}
