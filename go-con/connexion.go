package main

import (
	"context"
	"fmt"
	"os"

	pgx "github.com/jackc/pgx"
)

func main() {

	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connection to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())
	sum := 0
	rows, err := conn.Query(context.Background(), "select * from faculty;")
	//fmt.Println(tag, err)
	for rows.Next() {
		var n int
		var dn string
		var ln string
		err = rows.Scan(&n, &dn, &ln)
		if err != nil {
			fmt.Println("error occurred")
		}
		fmt.Println("ID no:", n, dn, ln)
		sum += n
	}
	fmt.Println(sum)

}
