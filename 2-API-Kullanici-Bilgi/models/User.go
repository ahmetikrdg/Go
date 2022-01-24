package models

type User struct {
	ID        int
	Username  string
	FirstName string
	LastName  string
	Interest  []string
	Profile   string
}
