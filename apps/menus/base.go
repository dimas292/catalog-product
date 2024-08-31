package menus

import (
	"database/sql"

	"github.com/go-chi/chi/v5"
)

func Run(router chi.Router, db *sql.DB){


	repo := newRepository(db)
	svc := newService(repo)
	handler := newHandler(svc)

	_ = handler
	
	
	router.Route("/api/menus", func(r chi.Router) {
		
	})
}