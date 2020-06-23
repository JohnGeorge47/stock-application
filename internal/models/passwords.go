package models

import (
	"errors"
	"fmt"
	"github.com/JohnGeorge47/stock-application/pkg/sql"
)

type UserPassword struct {
	UserId     string `json:"user_id"`
	UserEmail  string `json:"user_email"`
	Password   string `json:"password"`
	Updated_at string `json:"updated_at""`
}

func PasswordCreate(password UserPassword)error{
	fmt.Println(password)
	query := "INSERT INTO passwords(email_id,password,updated_at,user_id)values(?,?,?,(SELECT user_id FROM users WHERE email_id=?))"
	conn := sql.Connmanager
	fmt.Println(query)
	fmt.Println(password)
	val, err := conn.Insert(query, password.UserEmail, password.Password, password.Updated_at, password.UserEmail)
	fmt.Println(val)
	if err != nil {
		return err
	}
	return nil
}

func GetPassword(email string) (*string, error) {
	var pass string
	query := "SELECT password FROM passwords WHERE email_id=?"
	conn := sql.Connmanager
	rows, err := conn.Select(query, email)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		rows.Scan(&pass)
	}
	if pass == "" {
		return nil, errors.New("The email_id is not registered")
	}
	return &pass, err
}