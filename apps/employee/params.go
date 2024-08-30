package employee

type renderWeb struct {
	Title string
	Data  interface{}
}

type listEmployeeResponse struct {
	Id        int
	Name      string
	Address   string
	NIP       string
	CreatedAt string
}

type createNewEmpoyeesRequest struct {
	Name    string
	NIP     string
	Address string
}

