package main

import (
	"log"
	"os"
	"vue-api/cmd/api"
	"vue-api/internal/driver"
)

func main() {
	dsn := os.Getenv("DSN")

	db, err := driver.ConnectPostgres(dsn)
	if err != nil {
		log.Fatal("Cannot connect to database")
	}
	defer db.SQL.Close()

	app := api.NewApplication(8082, db.SQL)

	err = app.Serve()
	if err != nil {
		log.Fatal(err)
	}

}
