package models

import "database/sql"

type Database interface {
	ConnectionString() string

	Open	() 											(error)
	Ping	() 											(error)
	Close	() 											(error)
	Query 	(query string, args ...interface{}) 		(*sql.Rows, error)
	Exec 	(query string, args ...interface{}) 		(sql.Result, error)
	ExecTx 	(queries []string, args [][]interface{}) 	(error)
}