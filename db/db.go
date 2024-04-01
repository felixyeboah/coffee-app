package db

import (
	"database/sql"
	"fmt"
	"time"
)

type DB struct {
	DB *sql.DB
}

var dbConnection = &DB{}

const maxConnections = 10
const maxIdleConnections = 5
const maxConnectionLifetime = 5 * time.Minute

func connectPostgres(dsn string) (*DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(maxConnections)
	db.SetMaxIdleConns(maxIdleConnections)
	db.SetConnMaxLifetime(maxConnectionLifetime)

	err = testDBConnection(db)

	if err != nil {
		return nil, err
	}

	dbConnection.DB = db

	return dbConnection, nil
}

func testDBConnection(db *sql.DB) error {
	err := db.Ping()
	if err != nil {
		fmt.Println("Error connecting to the database", err)
		return err
	}

	fmt.Println("Successfully connected to the database")
	return nil
}
