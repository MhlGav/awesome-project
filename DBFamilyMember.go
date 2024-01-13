package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"os"
)

func DBFamily() {
	urlExample := "postgres://awesome_project:123@localhost:5436/test_db"
	conn, err := pgx.Connect(context.Background(), urlExample)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	var members string

	rows, err := conn.Query(context.Background(), "select members from family")
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("\nFAMILY:")
	for rows.Next() {
		rows.Scan(&members)
		fmt.Println(members)
	}
}
func DBPet() {
	urlExample := "postgres://awesome_project:123@localhost:5436/test_db"
	conn, err := pgx.Connect(context.Background(), urlExample)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	var name string

	rows, err := conn.Query(context.Background(), "select name from pet")
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("\nPET:")
	for rows.Next() {
		rows.Scan(&name)
		fmt.Println(name, "\n")
	}
}
