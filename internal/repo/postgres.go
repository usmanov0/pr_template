package repo

import (
	"database/sql"
	"fmt"
)

const (
	host     = "db"
	user     = "postgres"
	password = "root"
	dbName   = "dvdprogram"
	sslMode  = "disable"
)

func DatabaseConnection() (*sql.DB, error) {
	dbInfo := fmt.Sprint("host=$1  user=$2 password=$3 dnName=$4 sslMode=$5", host, user, password, dbName, sslMode)
	db, err := sql.Open("postgres", dbInfo)
	return db, err
}
