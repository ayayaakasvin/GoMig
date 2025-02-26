package parsing

import (
	"flag"
	"log"
	"migrationtool/internal/config"
)

func ParseFlags() *config.DatabaseConfig {
	host := flag.String("host", config.DefaultHost, "Host of the database. If not specified, default value would be \"Localhost\"")
	port := flag.Int("port", config.DefaultPort, "Port of the database. if not specified, default value would be 5432")
	user := flag.String("user", config.DefaultUser, "Username that will be used to open connection to the database. if not specified, default value would be postgres")
	password := flag.String("password", config.DefaultPassword, "Password for the database. if not specified, default value would be an empty string")
	dbname := flag.String("dbname", config.DefaultDBName, "Name of the database. if not specified, default value would be postgres")
	sslmode := flag.String("sslmode", config.DefaultSSLMode, "SSL mode for the database. if not specified, default value would be disable")

	flag.Parse()

	dbconfig := config.New(*host, *port, *user, *password, *dbname, *sslmode)

	log.Printf("Host: %s, Port: %d, User: %s, Password: %s, DBName: %s, SSLMode: %s", *host, *port, *user, *password, *dbname, *sslmode)

	return dbconfig
}