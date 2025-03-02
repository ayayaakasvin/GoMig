package scripts

import "errors"

var (
	EmptyMigrationConfig 	error = errors.New("migration configuration is empty")
	InvalidMigrationDir  	error = errors.New("migration directory is invalid")
	EmptyDatabaseConfig  	error = errors.New("database configuration is empty")
	NoSuchFile          	error = errors.New("file does not exist")
	NoSuchDirection     	error = errors.New("migration direction does not exist")
	DirectionNoFile     	error = errors.New("no migration files found for specified direction")
	UnknownMigration    	error = errors.New("unknown migration type")
	EmptySqlFile        	error = errors.New("SQL file is empty")
	EmptyDirectory      	error = errors.New("directory is empty")
)
