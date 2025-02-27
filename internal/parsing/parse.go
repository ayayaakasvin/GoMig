package parsing

import (
	"flag"
	"log"
	"migrationtool/internal/config/databaseconfig"
	"migrationtool/internal/config/migrationconfig"
)

func ParseFlags() (*databaseconfig.DatabaseConfig, *migrationconfig.MigrationConfig) {
	host := flag.String("host", databaseconfig.DefaultHost, "Host of the database. If not specified, default value would be \"Localhost\"")
	port := flag.Int("port", databaseconfig.DefaultPort, "Port of the database. if not specified, default value would be 5432")
	user := flag.String("user", databaseconfig.DefaultUser, "Username that will be used to open connection to the database. if not specified, default value would be postgres")
	password := flag.String("password", databaseconfig.DefaultPassword, "Password for the database. if not specified, default value would be an empty string")
	dbname := flag.String("dbname", databaseconfig.DefaultDBName, "Name of the database. if not specified, default value would be postgres")
	sslmode := flag.String("sslmode", databaseconfig.DefaultSSLMode, "SSL mode for the database. if not specified, default value would be disable")

	up := flag.Bool("up", false, "Set this flag to run the migration up")
	down := flag.Bool("down", false, "Set this flag to run the migration down")
	sourcePath := flag.String("path", "", "Path to the migration directory or file")

	flag.Parse()

	if *up && *down {
		log.Fatal("Cannot set both 'up' and 'down' flags")
	}

	if !*up && !*down {
		log.Fatal("Must set either 'up' or 'down' flag")
	}

	if *sourcePath == "" {
		log.Printf("path is not specified, would be set to \"./\"")
		*sourcePath = "./"
	}

	dbconfig := databaseconfig.New(*host, *port, *user, *password, *dbname, *sslmode)

	mtype := migrationconfig.Down
	if *up {
		mtype = migrationconfig.Up
	}
	migrconf := migrationconfig.New(migrationconfig.MigrationType(mtype), *sourcePath)
	

	log.Printf("Host: %s, Port: %d, User: %s, Password: %s, DBName: %s, SSLMode: %s", *host, *port, *user, *password, *dbname, *sslmode)
	log.Printf("Migration Path: %s, Up: %t, Down: %t", *sourcePath, *up, *down)

	return dbconfig, migrconf
}
