package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

// App description
type App struct {
	Router *mux.Router
	DB     *sql.DB
}

// Initialize description
func (a *App) Initialize(user, password, dbname, host string) {
	connectionString := fmt.Sprintf("user=%s password=%s dbname=%s host=%s sslmode=disable", user, password, dbname, host)

	var err error
	a.DB, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	a.Router = mux.NewRouter()
}

// Run description
func (a *App) Run(addr string) {}
