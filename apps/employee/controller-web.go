package employee

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path"
)

type handler struct {
	svc service
}

func newHandler(svc service) handler {
	return handler{svc: svc}
}

func (h handler) index(rw http.ResponseWriter, r *http.Request) {

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
		Data:  emp,
	}

	err = tmpl.Execute(rw, resp)
	if err != nil {
		log.Println("err")
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h handler) formCreateEmployee(rw http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(path.Join("external/public", "pages/employee/add.html"), path.Join("external/public", "layout/layout.html"))
	if err != nil {
		log.Println(err)
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	resp := renderWeb{
		Title: "Halaman Employee",
	}

	err = tmpl.Execute(rw, resp)
	if err != nil {
		log.Println(err)
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}
}

// lalu tambahkan method untuk process create employee
func (h handler) createEmployee(rw http.ResponseWriter, r *http.Request) {
	var req = createNewEmpoyeesRequest{}

	req = createNewEmpoyeesRequest{
		NIP:     r.FormValue("nip"),
		Address: r.FormValue("address"),
		Name:    r.FormValue("name"),
	}

	msg := ""
	if err := h.svc.createNewEmployee(r.Context(), req); err != nil {
		msg = `
			<script>
				alert("Tambah data pegawai gagal ! Error : %v")
				window.location.href="/employees"
			</script>
		`

		msg = fmt.Sprintf(msg, err.Error())
	} else {
		msg = `
			<script>
				alert("Tambah data pegawai berhasil !")
				window.location.href="/employees"
			</script>
		`

	}

	rw.Write([]byte(msg))
}

func (h handler) deleteEmployee(rw http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(rw, "Missing ID", http.StatusBadRequest)
		return
	}
	msg := ""
	if err := h.svc.deleteEmployeeByID(r.Context(), id); err != nil {
		log.Println("Gagal menghapus karyawan:", err)
		msg = fmt.Sprintf(msg, err.Error())
		return
	} else {
		msg = `
	<script>
		alert("hapus data berhasil !")
		window.location.href="/employees"
	</script>
	`
	}

	rw.Write([]byte(msg))

	http.Redirect(rw, r, "/employees", http.StatusSeeOther)
}
