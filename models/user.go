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
	users  []*User
	nextID = 1
)

/*
GetUsers function will get all the users saved.
*/
func GetUsers() []*User {
	return users
}

/*
AddUser will add a new user.
*/
func AddUser(u User) (User, error) {
	u.ID = nextID
	nextID++
	users = append(users, &u)
	return u, nil
}
