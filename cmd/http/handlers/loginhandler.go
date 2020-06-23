package handlers

import (
	"errors"
	"fmt"
	"net/http"
)

type Login struct {
	userEmail string
	password  string
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		contentType := r.Header.Get("Content-type")
		if contentType != "application/x-www-form-urlencoded" {
			w.WriteHeader(http.StatusUnsupportedMediaType)
			return
		}
		form, err := ValidateForm(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		//err=password.Login(form.userEmail, form.password)
		if err!=nil{
			http.Error(w,"Incorrect email_id or password",http.StatusForbidden)
		}
		fmt.Println(form)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func ValidateForm(r *http.Request) (*Login, error) {
	var loginform Login
	err := r.ParseForm()
	if err != nil {
		return nil, err
	}
	val, ok := r.PostForm["password"]
	if ok && len(val) != 0 {
		if val[0] == "" {
			return nil, errors.New("Password field cannot be empty")
		}
		loginform.password = val[0]
	} else {
		return nil, errors.New("Missing Password field")
	}
	val, ok = r.PostForm["user_email"]
	if ok && len(val) != 0 {
		if val[0] == "" {
			return nil, errors.New("user_email field cannot be empty")
		}
		loginform.userEmail = val[0]
	} else {
		return nil, errors.New("Missing user_email field")
	}
	return &loginform, nil
}