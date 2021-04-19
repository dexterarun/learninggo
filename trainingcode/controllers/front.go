package controllers

import "net/http"

func RegisterControllers() {
	uc := newUserController()

	//route 1 - to handle requests to controller home /users
	http.Handle("/users", *uc)
	//route 2 - to handle requests like /users/5 etc.
	http.Handle("/users/", *uc)
}
