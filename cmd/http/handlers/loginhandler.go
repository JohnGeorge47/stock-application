package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/JohnGeorge47/stock-application/internal/core/user"
	"net/http"
)

type Login struct {
	userEmail string
	password  string
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	EnableCors(&w)
	switch r.Method {
	case "POST":
		var loginresp SignupResponse
		loginresp.Method = http.MethodPost
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
		if err != nil {
			http.Error(w, "Incorrect email_id or password", http.StatusForbidden)
		}
		resp, err := user.Login(form.userEmail, form.password)
		if err != nil {
			fmt.Println(err)
			http.Error(w, "Username or password is incorrect", http.StatusBadRequest)
			return
		}
		loginresp.Success = true
		loginresp.RequestToken = *resp
		jsonresp, err := json.Marshal(loginresp)
		if err != nil {
			fmt.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonresp)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func EnableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
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
	val, ok = r.PostForm["email_id"]
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
