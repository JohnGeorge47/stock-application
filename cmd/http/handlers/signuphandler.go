package handlers

import (
	"errors"
	"net/http"
	"github.com/JohnGeorge47/stock-application/internal/core/createuser"
)





func SignupHandler(w http.ResponseWriter,r *http.Request){
	switch r.Method {
	case "POST":
		contentType := r.Header.Get("Content-type")
		if contentType != "application/x-www-form-urlencoded" {
			w.WriteHeader(http.StatusUnsupportedMediaType)
			return
		}
		user,err:=ValidateSignupForm(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		err=createuser.Create(*user)
		if err!=nil{
			http.Error(w,err.Error(),http.StatusInternalServerError)
			return
		}
	}
}

func ValidateSignupForm(r *http.Request) (*createuser.Signup, error) {
	var user createuser.Signup
	err:=r.ParseForm()
	if err != nil {
		return nil, err
	}

	val, ok := r.PostForm["password"]
	if ok && len(val) != 0 {
		if val[0] == "" {
			return nil, errors.New("Password field cannot be empty")
		}
		user.Password = val[0]
	} else {
		return nil, errors.New("Missing Password field")
	}
	val, ok = r.PostForm["user_email"]
	if ok && len(val) != 0 {
		if val[0] == "" {
			return nil, errors.New("user_email field cannot be empty")
		}
		user.EmailId = val[0]
	} else {
		return nil, errors.New("Missing user_email field")
	}
	return &user, nil
}