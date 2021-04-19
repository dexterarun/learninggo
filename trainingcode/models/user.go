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

func GetUsers() []*User {
	return users
}

func AddUser(u User) (User, error) {
	u.ID = nextID
	nextID++
	users = append(users, &u)
	return u, nil // user=u, error=nil are the two return values for (User, error)
}
