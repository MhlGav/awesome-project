package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"os"
)

func DBDescription() {
	urlExample := "postgres://awesome_project:123@localhost:5436/test_db"
	conn, err := pgx.Connect(context.Background(), urlExample)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	var parameters string
	var values int64

	//err = conn.QueryRow(context.Background(), "select parameters, values from description").Scan(&parameters, &values)

	rows, err := conn.Query(context.Background(), "select parameters, values from description")

	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("DESCRIPTION:")
	for rows.Next() {
		rows.Scan(&parameters, &values)
		fmt.Println(parameters, values)
	}
}

func DBFurniture() {
	urlExample := "postgres://awesome_project:123@localhost:5436/test_db"
	conn, err := pgx.Connect(context.Background(), urlExample)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	var bedroom string
	var living_room string
	var kitchen string

	rows, err := conn.Query(context.Background(), "select bedroom, living_room, kitchen from furniture")
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("\nFURNITURA:")
	for rows.Next() {
		rows.Scan(&bedroom, &living_room, &kitchen)
		fmt.Println("bedroom:", bedroom, "\n living_room:", living_room, "\n  kitchen:", kitchen)
	}
}
