package parsing

import (
    "flag"
    "log"
	"os"
	
    "migrationtool/internal/config/databaseconfig"
    "migrationtool/internal/config/migrationconfig"
)

// ParseFlags processes command-line flags to configure database connection and migration settings.
// ParseFlags processes command-line flags to configure database and migration settings.
// It handles both database connection parameters and migration control flags.
//
// Database configuration flags include:
//   - host: Database host address (default: "localhost")
//   - port: Database port number (default: 5432)
//   - user: Database username (default: "postgres")
//   - password: Database password (default: "")
//   - dbname: Database name (default: "postgres")
//   - sslmode: SSL mode for connection (default: "disable")
//
// Migration control flags include:
//   - up: Flag to perform migration up
//   - down: Flag to perform migration down
//   - path: Directory path containing migration files (default: "./")
//
// The function validates that exactly one of 'up' or 'down' flags is set.
// If path is not specified, it defaults to the current directory.
//
// Returns:
//   - *databaseconfig.DatabaseConfig: Configuration for database connection
//   - *migrationconfig.MigrationConfig: Configuration for migration operation
//
// The function will terminate with log.Fatal if invalid flag combinations are detected.
func ParseFlags() (*databaseconfig.DatabaseConfig, *migrationconfig.MigrationConfig) {
    fs := flag.NewFlagSet("migrationtool", flag.ExitOnError)

    // Database flags
    host := fs.String("host", databaseconfig.DefaultHost, "Host of the database. If not specified, default value would be \"Localhost\"")
    port := fs.Int("port", databaseconfig.DefaultPort, "Port of the database. if not specified, default value would be 5432")
    user := fs.String("user", databaseconfig.DefaultUser, "Username that will be used to open connection to the database. if not specified, default value would be postgres")
    password := fs.String("password", databaseconfig.DefaultPassword, "Password for the database. if not specified, default value would be an empty string")
    dbname := fs.String("dbname", databaseconfig.DefaultDBName, "Name of the database. if not specified, default value would be postgres")
    sslmode := fs.String("sslmode", databaseconfig.DefaultSSLMode, "SSL mode for the database. if not specified, default value would be disable")

    // Migration flags
    up := fs.Bool("up", false, "Set this flag to run the migration up")
    down := fs.Bool("down", false, "Set this flag to run the migration down")
    sourcePath := fs.String("path", "", "Path to the migration directory or file")

    // Parse the flags
    fs.Parse(os.Args[1:])

    // Validate migration flags
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