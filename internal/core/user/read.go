package user

import (
	"github.com/JohnGeorge47/stock-application/internal/models"
	"golang.org/x/crypto/bcrypt"
)

func Login(email string, password string) error {
	hashpwd, err := models.GetPassword(email)
	if err != nil {
		return err
	}
	err = bcrypt.CompareHashAndPassword([]byte(*hashpwd), []byte(password))
	if err != nil {
		return err
	}
	return nil
}

func Validate(session_token string, email_id string) (*bool, error) {
	val, err := models.CheckSession(email_id, session_token)
	if err != nil {
		return nil, err
	}
	return val, nil
}
