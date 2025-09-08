package main

import (
	"log"
	"os"

	"github.com/ayayaakasvin/GoMig/internal/models/postgresql"
	"github.com/ayayaakasvin/GoMig/internal/parsing"
	"github.com/ayayaakasvin/GoMig/internal/scripts"
)

func main() {
	dbconf, migrconf := parsing.ParseFlags()
	db := postgresql.New(dbconf)
	defer db.Close()
	
	log.Printf("Database configuration: %+v", dbconf)
    log.Printf("Migration configuration: %+v", migrconf)


	sqlScripts, err := scripts.ParseMigrationFiles(migrconf)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Number of SQL scripts: %d", len(sqlScripts))

	err = scripts.ExecuteScripts(db,sqlScripts)
	if err != nil {
		log.Fatalf("Failed to execute scripts: %v", err)
	}

	log.Println("Migration completed successfully")
	os.Exit(0)
}
