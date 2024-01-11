package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"os"
)

func DBbuilding() {
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

	defer conn.Close(context.Background())

	var bedroom string
	var living_room string
	var kitchen string

	rows1, err := conn.Query(context.Background(), "select bedroom, living_room, kitchen from furniture")
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("\nFURNITURA:")
	for rows1.Next() {
		rows1.Scan(&bedroom, &living_room, &kitchen)
		fmt.Println("bedroom:", bedroom, "\n living_room:", living_room, "\n  kitchen:", kitchen)
	}

	defer conn.Close(context.Background())

	var members string

	rows2, err := conn.Query(context.Background(), "select members from family")
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("\nFAMILY:")
	for rows2.Next() {
		rows2.Scan(&members)
		fmt.Println(members)
	}

	defer conn.Close(context.Background())

	var name string

	rows3, err := conn.Query(context.Background(), "select name from pet")
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("\nPET:")
	for rows3.Next() {
		rows3.Scan(&name)
		fmt.Println(name, "\n")
	}
}
