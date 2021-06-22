package dto

type Users struct {
	Id       int
	Name     string
	Username string
	Email    string
	Phone    string
	Website  string
	Address  interface{}
	Company  interface{}
}
