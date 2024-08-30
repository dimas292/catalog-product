package apps

import (
	"catalog-product/apps/employee"
	"database/sql"
	"html/template"
	"log"
	"net/http"
	"path"

	"github.com/go-chi/chi/v5"
)


func Run(appPort string, db *sql.DB){

	router := chi.NewRouter()

	registerRouting(router, db)

	log.Printf("server running at port %v\n", appPort)
	http.ListenAndServe(appPort, router)
}

func registerRouting(router chi.Router, db *sql.DB){
	router.Get("/", getIndex)

	employee.Run(router, db)

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