package postgresql

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"

	"migrationtool/internal/config/databaseconfig"
)

// PostgresSql represents the configuration required to connect to a PostgreSQL database.
type PostgresSql struct {
	Config *databaseconfig.DatabaseConfig // Config holds the database configuration.
	DB     *sql.DB                // DB is the sql.DB object for database connection.
}

// ConnectionString returns the connection string for the PostgreSQL database.
func (ps *PostgresSql) ConnectionString() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s", ps.Config.Host, ps.Config.Port, ps.Config.User, ps.Config.Password, ps.Config.DBName, ps.Config.SSLMode)
}

// New creates a new instance of PostgresSql with the given configuration parameters. In case of error, it log.Fatal, with error value.
func New(config *databaseconfig.DatabaseConfig) (*PostgresSql) {
	db := &PostgresSql{
		Config: config,
	}

	if err := db.Open(); err != nil {
		log.Fatal(err)
	}

	if err := db.Ping();err != nil {
		log.Fatal(err)
	}

	return db
}

// Open opens a new database connection and assigns it to the DB field.
func (ps *PostgresSql) Open() error {
	connStr := ps.ConnectionString()
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return err
	}
	ps.DB = db
	return nil
}

// Close closes the database connection.
func (ps *PostgresSql) Close() error {
	if ps.DB != nil {
		return ps.DB.Close()
	}
	return nil
}

// Ping verifies a connection to the database is still alive, establishing a connection if necessary.
func (ps *PostgresSql) Ping() error {
	if ps.DB == nil {
		if err := ps.Open(); err != nil {
			return err
		}
	}
	return ps.DB.Ping()
}

// Query executes a query that returns rows, typically a SELECT.
func (ps *PostgresSql) Query(query string, args ...interface{}) (*sql.Rows, error) {
	if ps.DB == nil {
		if err := ps.Open(); err != nil {
			return nil, err
		}
	}
	return ps.DB.Query(query, args...)
}

// Exec executes a query without returning any rows, typically an INSERT, UPDATE, or DELETE.
func (ps *PostgresSql) Exec(query string, args ...interface{}) (sql.Result, error) {
	if ps.DB == nil {
		if err := ps.Open(); err != nil {
			return nil, err
		}
	}
	return ps.DB.Exec(query, args...)
}

// ExecTx executes a series of queries within a transaction.
func (ps *PostgresSql) ExecTx(queries []string, args [][]interface{}) error {
	if ps.DB == nil {
		if err := ps.Open(); err != nil {
			return err
		}
	}

	tx, err := ps.DB.Begin()
	if err != nil {
		return err
	}

	for i, query := range queries {
		if _, err := tx.Exec(query, args[i]...); err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit()
}
