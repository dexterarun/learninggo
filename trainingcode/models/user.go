package models

type User struct {
	ID        int
	FirstName string
	LastName  string
}

var (
	//pointer to slice of user objects.
	users  []*User
	nextID = 1 //within var scope dont need := , compiler knows it's an integer type(int32). only time data type would be mentioned is if variable were to be used that's not going to be the default.
)
