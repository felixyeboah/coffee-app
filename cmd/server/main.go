package main

import (
	"coffee-app/db"
	"coffee-app/routes"
	"coffee-app/services"
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

type Config struct {
	Port string
}

type Application struct {
	Config Config
	Models services.Models
}

func (app *Application) Run() error {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = app.Config.Port
	}

	fmt.Println("Server is running on port", port)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: routes.Routes(),
	}

	err = server.ListenAndServe()
	if err != nil {
		return err
	}

	return server.ListenAndServe()
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// setting up configuration
	config := Config{
		Port: os.Getenv("PORT"),
	}

	// connecting database
	dsn := os.Getenv("DSN")
	if dsn == "" {
		log.Fatal("DSN is required")
	}
	dbConnection, err := db.ConnectPostgres(dsn)
	if err != nil {
		log.Fatal("Error connecting to the database", err)
	}
	defer func(DB *sql.DB) {
		err := DB.Close()
		if err != nil {
			log.Fatal("Error closing the database connection", err)
		}
	}(dbConnection.DB)

	// running the application
	app := &Application{
		Config: config,
		Models: services.New(dbConnection.DB),
	}
	err = app.Run()
	if err != nil {
		log.Fatal(err)
	}
}
