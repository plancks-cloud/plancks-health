package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"git.amabanana.com/plancks-cloud/pc-go-brutus/controller/db"
)

func Health(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(db.HealthCheck())
}

//CORSHandler does CORS check
func CORSHandler(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-type")
		w.Header().Set("Allow", "GET,POST,OPTIONS")

		if r.Method == "OPTIONS" {
			fmt.Fprintf(w, "Options")
		} else {
			f.ServeHTTP(w, r)
		}
	}
}
