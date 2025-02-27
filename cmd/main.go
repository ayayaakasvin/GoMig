package main

import (
	"log"
	"os"
	
	"migrationtool/internal/models/postgresql"
	"migrationtool/internal/parsing"
)

func main() {
	dbconf, migrconf := parsing.ParseFlags()
	db := postgresql.New(dbconf)
	log.Print(db, migrconf)

	// TODO: parse the migration files and execute them
	os.Exit(0)
}
