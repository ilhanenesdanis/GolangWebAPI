package models

type Users struct {
	ID        int
	UserName  string
	FirstName string
	LastName  string
	Profile string
	Interest  []Interest
}
