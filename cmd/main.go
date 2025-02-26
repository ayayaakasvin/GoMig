package main

import (
	"migrationtool/internal/models/postgresql"
	"migrationtool/internal/parsing"
	"os"
)

func main() {
	config := parsing.ParseFlags()
	db := postgresql.New(config)

	// TODO: parse the migration files and execute them
	os.Exit(0)
}