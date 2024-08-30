// file : apps/employee/base.go

package employee

import (
	"database/sql"
	"net/http"

	"github.com/go-chi/chi/v5"
)

// method ini akan dipanggil oleh apps/server.go
// method ini sebagai entry point ke base modules
// base modules membutuhkan chi.Router dan sql.DB
func Run(router chi.Router, db *sql.DB) {
    // init repository dengan memasukkan object database didalamnya
	repo := newRepository(db)

    // init service dengan memasukkan object repository didalamnya
	svc := newService(repo)

    // init handler dengan memasukkan object service didalamnya
	handler := newHandler(svc)

	// ini berfungsi menghandle sebuah folder yg isinya adalah
	// static file. Biasanya ini bersifat asset asset
	fileServer := http.FileServer(http.Dir("external/public/assets"))
    
    // ini berfungsi untuk open ke public jika ada asset yang akan kita gunakan
	router.Handle("/public/", http.StripPrefix("/external/public", fileServer))
  
    // proses grouping routing
    // jadi routing didalamnya akan diawali dengan /employees
	router.Route("/employees", func(r chi.Router) {
        // endpoint : GET /employees
		r.Get("/", handler.index)
		r.Post("/process/add", handler.createEmployee)
		r.Get("/add", handler.formCreateEmployee)

	})
}
