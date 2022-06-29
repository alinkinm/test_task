package main

import (
	"database/sql"
	"fmt"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres504"
	dbname   = "postgres"
)

var psqlconn = fmt.Sprintf("host = %s port = %d user = %s password = %s dbname = %s "+
	"sslmode=disable", host, port, user, password, dbname)

func connectToDataBase() *sql.DB {
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)
	return db
}
