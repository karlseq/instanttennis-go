package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var (
	host     = "database-test.c3lf9o6zdmvb.us-east-1.rds.amazonaws.com"
	password = "Karloo123!"
	port     = "3306"
	user     = "root"
	dbName   = "InstantTennis"
)

func GetConnection() (*sql.DB, error) {
	return sql.Open("mysql", dsn())

}

func dsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, password, host, port, dbName)
}
