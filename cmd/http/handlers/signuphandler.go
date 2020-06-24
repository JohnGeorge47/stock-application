package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/JohnGeorge47/stock-application/internal/core/user"
	"net/http"
)

type SignupResponse struct {
	Method       string `json:"method"`
	Success      bool   `json:"success"`
	RequestToken string `json:"request_token"`
}

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		var sresponse SignupResponse
		sresponse.Method = r.Method
		contentType := r.Header.Get("Content-type")
		if contentType != "application/x-www-form-urlencoded" {
			w.WriteHeader(http.StatusUnsupportedMediaType)
			return
		}
		signup, err := ValidateSignupForm(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		reqToken, err := user.Create(*signup)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Println(*reqToken)
		sresponse.Success = true
		sresponse.RequestToken = *reqToken
		jsonres, err := json.Marshal(sresponse)
		fmt.Println(err)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonres)

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func ValidateSignupForm(r *http.Request) (*user.Signup, error) {
	var signup user.Signup
	err := r.ParseForm()
	if err != nil {
		return nil, err
	}

	val, ok := r.PostForm["password"]
	if ok && len(val) != 0 {
		if val[0] == "" {
			return nil, errors.New("Password field cannot be empty")
		}
		signup.Password = val[0]
	} else {
		return nil, errors.New("Missing Password field")
	}
	val, ok = r.PostForm["user_email"]
	if ok && len(val) != 0 {
		if val[0] == "" {
			return nil, errors.New("user_email field cannot be empty")
		}
		signup.EmailId = val[0]
	} else {
		return nil, errors.New("Missing user_email field")
	}
	val, ok = r.PostForm["user_name"]
	if ok && len(val) != 0 {
		if val[0] == "" {
			return nil, errors.New("user_email field cannot be empty")
		}
		signup.UserName = val[0]
	} else {
		return nil, errors.New("Missing user_email field")
	}

	return &signup, nil
}
