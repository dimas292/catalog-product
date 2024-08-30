package employee

import (
	"html/template"
	"log"
	"net/http"
	"path"
	"fmt"
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

// file : /apps/employee/controller-web.go


// ... kode sebelumnya

// tambahkan method untuk menampilkan form
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

	// proses render file yang telah kita panggil diatas.
	// method Execute membutuhkan 2 parameter, yaitu sebuah ResponseWriter dan sebuah data.
	// karnea pada method Index ini kita tidak membutuhkan data, maka cukup ditulis dengan nil
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

	// proses pengambilan nilai input dari html.
	// nip, address, dan name itu akan didapat dari tag input di properti `name` pada html
	// contohnya : <input name="nip" />
	// atau <input name="address" />
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

