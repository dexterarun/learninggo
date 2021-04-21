package controllers

import (
	"encoding/json"
	"net/http"
	"regexp"
	"strconv"

	"github.com/pluralsight/webservice/models"
)

//package name must match directory we're in.

type userController struct {
	userIDPattern *regexp.Regexp
}

//function changes to a method when you specify the type that you want to bind the function to.
// here uc is bound to the user controller.
// ServeHTTP is our method name and it requires two parameters - responsewriter and request object.
func (uc userController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//w.Write([]byte("Hello from User Controller!"))
	// byte slice is an alias for a string. conversely a string is an alias for a byte slice. above is syntax for string to byteslice.
	//go lang dox:
	// type Handler interface {
	//ServeHTTP(ResponseWriter, *Request)
	// }

		if r.URL.Path == "/users" {
			switch r.Method {
			case http.MethodGet : uc.getAll(w, r)
			case http.MethodPost : uc.post(w, r)
			default: w.WriteHeader(http.StatusNotImplemented)
			}
		} else {
				matches := uc.userIDPattern.FindAllStringSubmatch(r.URL.Path)
				if len(matches) == 0 {
					w.WriteHeader(http.StatusNotFound)
				}
				id,err := strconv.Atoi(matches[1])  //subgroup match containing the id value
				if err != nil {
					w.WriteHeader(http.StatusNotFound)
				}
				switch r.Method {
				case http.MethodGet:
					uc.get(id, w)
				case http.MethodPut:
					uc.put(id, w, r)
				case http.MethodDelete:
					uc.delete(id, w)
				default: 
				w.WriteHeader(http.StatusNotFound)
				}

			}
		}

}

func (uc *userController) get(id int, w http.ResponseWriter) {
	u, err := models.GetUserByID(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	encodeResponseAsJSON(u, w)
}

func (uc *userController) getAll(w http.ResponseWriter, r *http.Request) {
	encodeResponseAsJSON(models.GetUsers(), w)
}

func (uc *userController) post(w http.ResponseWriter, r *http.Request) {

	u, err := uc.parseRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Could not parse user object"))
		return
	}
	u, err = models.AddUser(u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	encodeResponseAsJSON(u, w)
}

func (uc *userController) put(id int, w http.ResponseWriter, r *http.Request) {
	u, err := uc.parseRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Could not parse user object"))
		return
	}
	if id != u.ID {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("ID of submitted user must match ID in URL"))
		return
	}

	u, err = models.UpdateUser(u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	encodeResponseAsJSON(u, w)
}

func (uc *userController) delete(id int, w http.ResponseWriter) {
	err := models.RemoveUserById(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (uc *userController) parseRequest(r *http.Request) (models.User, error) {
	dec := json.NewDecoder(r.Body)
	var u models.User
	err := dec.Decode(&u)
	if err != nil {
		return models.User{}, err
	}
	return u, nil
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
