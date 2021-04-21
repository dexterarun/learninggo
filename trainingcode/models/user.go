package models

import (
	"errors"
	"fmt"
)

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
	if u.ID != 0 {
		return User{}, errors.New("New user must not include ID or it must be set ") //returning empty user obj and an error.
	}
	u.ID = nextID
	nextID++
	users = append(users, &u)
	return u, nil // user=u, error=nil are the two return values for (User, error)
}

func GetUserByID(id int) (User, error) {
	for _, u := range users {
		if u.ID == id {
			return *u, nil
		}
	}
	return User{}, fmt.Errorf("User with ID '%v' not found", id)
}

func UpdateUser(u User) (User, error) {
	for i, candidate := range users {
		if candidate.ID == u.ID {
			users[i] = &u
			return u, nil
		}
	}
	return User{}, fmt.Errorf("User with ID '%v' not found", u.ID)
}

func RemoveUserById(id int) error {
	for i, u := range users {
		if u.ID == id {
			users = append(users[:i], users[i+1:]...) //performing a splice operation on a slice.
			return nil
		}
	}
	return fmt.Errorf("User with ID '%v' could not be found", id)
}
