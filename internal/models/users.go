package models

import (
	"fmt"
	"github.com/JohnGeorge47/stock-application/pkg/sql"
)

type User struct {
	EmailId    string `json:"email_id"`
	UserName   string `json:"user_name"`
	Updated_at string `json:"updated_at"`
}

func CreateUser(user User) (*int64, error) {
	query := "INSERT INTO users(email_id,user_name) VALUES(?,?)"
	conn := sql.Connmanager
	lastid, err := conn.Insert(query, user.EmailId, user.UserName)
	if err != nil {
		return nil, err
	}
	fmt.Println(lastid)
	return lastid, nil
}

func CreateSession(email_id string, session_token string) error {
	query := `INSERT INTO sessions(session_token,user_id)
             SELECT ?,user_id from users WHERE email_id=?`
	conn := sql.Connmanager
	_, err := conn.Insert(query, session_token, email_id)
	if err != nil {
		return err
	}
	return nil
}

func CheckSession(email_id string, session_token string) (*bool, error) {
	query := `SELECT EXISTS( SELECT * 
							FROM sessions s
							JOIN users u
							ON s.user_id=u.user_id
							WHERE u.email_id=?
							AND	  s.session_token=?
							)`
	conn := sql.Connmanager
	fmt.Println(email_id, session_token)
	res, err := conn.Select(query, email_id, session_token)
	defer res.Close()
	for res.Next() {
		var val bool
		if err := res.Scan(&val); err != nil {
			fmt.Println(err)
		}
		return &val, nil
	}
	if err != nil {
		return nil, err
	}
	fmt.Println(res)
	defaultval := false
	return &defaultval, nil
}
