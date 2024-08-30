package employee

type renderWeb struct{
	Title string
	Data interface{}
}

type listEmployeeResponse struct {
	Id        int
	Name      string
	Address   string
	NIP       string
	CreatedAt string

}
