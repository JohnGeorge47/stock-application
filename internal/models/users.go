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

func CreateUser(user User) (*int64,error) {
	query := "INSERT INTO users(email_id,user_name) VALUES(?,?)"
	conn := sql.Connmanager
	lastid, err := conn.Insert(query, user.EmailId, user.UserName)
	if err != nil {
		return nil,err
	}
	fmt.Println(lastid)
	return  lastid,nil
}

