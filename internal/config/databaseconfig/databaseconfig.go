package databaseconfig

type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	SSLMode  string
}

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
