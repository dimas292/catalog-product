package apps

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"
	"path"

	"github.com/go-chi/chi"
)


func Run(appPort string, db *sql.DB){

	router := chi.NewRouter()

	registerRouting(router)

	log.Printf("server running at port %v\n", appPort)
	http.ListenAndServe(appPort, router)
}

func registerRouting(router chi.Router){
	router.Get("/", getIndex)
}

func getIndex(rw http.ResponseWriter, r *http.Request){

	tmpl, err := template.ParseFiles(path.Join("external/public", "pages/home/index.html"), path.Join("external/public", "layout/layout.html"))
	if err != nil {
		log.Println(err)
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return 
	}


	err = tmpl.Execute(rw, nil)
	if err != nil {
		log.Println(err)
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return 
	}
}