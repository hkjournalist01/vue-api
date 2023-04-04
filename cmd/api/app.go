package api

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"vue-api/internal/data"
)

type config struct {
	port int
}

type application struct {
	config      config
	infoLog     *log.Logger
	errorLog    *log.Logger
	models      data.Models
	devices     data.Devices
	db          *sql.DB
	environment string
}

func NewApplication(port int, db *sql.DB) *application {
	var cfg config
	cfg.port = 8082

	infoLog := log.New(os.Stdout, "Info\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "Error\t", log.Ldate|log.Ltime|log.Lshortfile)

	environment := os.Getenv("ENV")

	app := &application{
		config:      cfg,
		infoLog:     infoLog,
		errorLog:    errorLog,
		models:      data.New(),
		db:          db,
		environment: environment,
	}

	return app
}

func (app *application) Serve() error {
	app.infoLog.Println("API listening on port", app.config.port)
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", app.config.port),
		Handler: app.routes(),
	}

	return srv.ListenAndServe()
}
