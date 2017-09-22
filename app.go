package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

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
	a.initializeRoutes()
}

// Run description
func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, a.Router))
}

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/servers", a.getServers).Methods("GET")
	a.Router.HandleFunc("/server/{id:[a-zA-Z0-9]+}", a.getServer).Methods("GET")
	a.Router.HandleFunc("/packages/{id:[a-zA-Z0-9]+}", a.getPackages).Methods("GET")
}

func (a *App) getServers(w http.ResponseWriter, r *http.Request) {
	servers, err := getMachineIDs(a.DB)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, servers)
}

func (a *App) getServer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	machineid := vars["id"]
	server, err := getSysInfoByMI(a.DB, machineid)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, server)
}

func (a *App) getPackages(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	machineid := vars["id"]
	server, err := getPackagesByMI(a.DB, machineid)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, server)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
