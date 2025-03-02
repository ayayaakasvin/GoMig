package databaseconfig

type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	SSLMode  string
}

// New creates and returns a new DatabaseConfig with the specified connection parameters.
// Parameters:
//   - host: The database server host address
//   - port: The database server port number
//   - user: The database user name
//   - password: The database user password
//   - dbname: The name of the database to connect to
//   - sslmode: The SSL mode for the connection (e.g., "disable", "require", "verify-full")
// Returns:
//   - *DatabaseConfig: A pointer to the newly created DatabaseConfig struct
func New(host string, port int, user, password, dbname, sslmode string) *DatabaseConfig {
	return &DatabaseConfig{
		Host:     host,
		Port:     port,
		User:     user,
		Password: password,
		DBName:   dbname,
		SSLMode:  sslmode,
	}
}

const (
	DefaultHost     = "localhost"
	DefaultPort     = 5432
	DefaultUser     = "postgres"
	DefaultPassword = ""
	DefaultDBName   = "postgres"
	DefaultSSLMode  = "disable"
)
