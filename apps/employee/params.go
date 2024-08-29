package employee

type renderWeb struct {
	Title string
	Data  interface{}
}

type listEmployeeResponse struct {
	Id        int
	Name      string
	NIP       string
	Address   string
	CreatedAt string
}
