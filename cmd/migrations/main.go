package main

import (
	"database/sql"
	"flag"
	"log"

	"github.com/pressly/goose"

	_ "github.com/lib/pq"
)

func main() {
	dir := flag.String("dir", "database/migrations/postgres", "migration path")
	dsn := flag.String("dsn", "postgres://user:pass@localhost:5435/inventory?sslmode=disable", "connection string")

	flag.Parse()

	if err := goose.SetDialect("postgres"); err != nil {
		log.Fatal(err)
	}

	db, err := sql.Open("postgres", *dsn)
	if err != nil {
		log.Fatalf("goose: failed to open DB: %v\n", err)
	}

	defer func() {
		if err := db.Close(); err != nil {
			log.Fatalf("goose: failed to close DB: %v\n", err)
		}
	}()

	// for now, just support one command
	command := "up"
	if err := goose.Run(command, db, *dir, ""); err != nil {
		log.Fatalf("goose %v: %v", command, err)
	}
}
