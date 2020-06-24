package user

import (
	"fmt"
	"github.com/JohnGeorge47/stock-application/internal/models"
	"github.com/JohnGeorge47/stock-application/pkg/uuid"
	"golang.org/x/crypto/bcrypt"
)

func Login(email string, password string) (*string, error) {
	hashpwd, err := models.GetPassword(email)
	if err != nil {
		return nil, err
	}
	fmt.Println(email, hashpwd)
	err = bcrypt.CompareHashAndPassword([]byte(*hashpwd), []byte(password))
	if err != nil {
		return nil, err
	}
	reqtoken := uuid.GetUUID()
	err = models.CreateSession(email, reqtoken)
	if err != nil {
		return nil, err
	}
	return &reqtoken, nil
}

func Validate(session_token string, email_id string) (*bool, error) {
	val, err := models.CheckSession(email_id, session_token)
	if err != nil {
		return nil, err
	}
	return val, nil
}
