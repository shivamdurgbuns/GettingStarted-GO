package models

/*
User is the structre of a user.
*/
type User struct {
	ID        int
	FirstName string
	LastName  string
}

var (
	user   []*User
	nextID = 1
)
