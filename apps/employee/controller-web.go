package employee

import (
	"html/template"
	"log"
	"net/http"
	"path"
)

type handler struct {
	svc service
}

func newHandler(svc service) handler{
	return handler{svc: svc}
}

func (h handler) index(rw http.ResponseWriter, r *http.Request){

	tmpl, err := template.ParseFiles(path.Join("external/public", "pages/employee/index.html"), path.Join("external/public", "layout/layout.html"))

	if err != nil {
		log.Println(err)
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return 
	}

	emp, err := h.svc.listEmployee(r.Context())
	if err != nil {
		log.Println(err)
	}

	resp := renderWeb{
		Title: "halaman web",
		Data: emp,
	}

	err = tmpl.Execute(rw, resp)
	if err != nil {
		log.Println("err")
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
}

