package controllers

import (
	"net/http"
	"regexp"
)

//package name must match directory we're in.

type userController struct {
	userIDPattern *regexp.Regexp
}

//function changes to a method when you specify the type that you want to bind the function to.
// here uc is bound to the user controller.
// ServeHTTP is our method name and it requires two parameters - responsewriter and request object.
func (uc userController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from User Controller!"))
	// byte slice is an alias for a string. conversely a string is an alias for a byte slice. above is syntax for string to byteslice.
	//go lang dox:
	// type Handler interface {
	//ServeHTTP(ResponseWriter, *Request)
	// }
}

//constructor function - start with new then name of controller. that's go convention.
func newUserController() *userController {
	//with structs we can take address of struct object immediately after init else this would be ringing alarm bells. but ok with structs.
	// also note that usercontroller here is a local variable for which we return the address out of the function. In c++ this would be a worry as anything could overwrite that memory after losing scope of method.
	// but in go, that variable is promoted up to whatever scope it needs to be at for its use.
	return &userController{
		userIDPattern: regexp.MustCompile(`^/users/(\d+)/?`),
	}
}
