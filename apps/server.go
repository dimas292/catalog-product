package apps

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

func Run(appPort string, db *sql.DB) {

	router := chi.NewRouter()

	log.Printf("server running at %v\n", appPort)
	if err := http.ListenAndServe(appPort, router); err != nil {
		panic(err)
	}

}
