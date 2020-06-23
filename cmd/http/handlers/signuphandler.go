package handlers

import (
	"errors"
	"fmt"
	"github.com/JohnGeorge47/stock-application/internal/core/user"
	"github.com/JohnGeorge47/stock-application/pkg/uuid"
	"net/http"
)



func SignupHandler(w http.ResponseWriter,r *http.Request){
	switch r.Method {
	case "POST":
		contentType := r.Header.Get("Content-type")
		if contentType != "application/x-www-form-urlencoded" {
			w.WriteHeader(http.StatusUnsupportedMediaType)
			return
		}
		signup,err:=ValidateSignupForm(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		err= user.Create(*signup)
		reqtoken:=uuid.GetUUID()
		fmt.Println(reqtoken)
		if err!=nil{
			http.Error(w,err.Error(),http.StatusInternalServerError)
			return
		}
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func ValidateSignupForm(r *http.Request) (*user.Signup, error) {
	var signup user.Signup
	err:=r.ParseForm()
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