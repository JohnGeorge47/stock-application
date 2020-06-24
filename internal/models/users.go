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

func GetLoggedinUsers() (*[]string, error) {
	query := `SELECT DISTINCT user_id from sessions`
	conn := sql.Connmanager
	rows, err := conn.Select(query)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	var user_id_list []string
	defer rows.Close()
	for rows.Next() {
		var val string
		if err := rows.Scan(&val); err != nil {
			fmt.Println(err)
		}
		user_id_list = append(user_id_list, val)
	}
	return &user_id_list, nil
}

type SubscribedStock struct {
	EmailId   string `json:"email_id"`
	StockName string `json:"stock_name"`
	Value     string `json:"value"`
}

func GetAllUserSubscribedStocks(user_email string) (*[]SubscribedStock, error) {

	query := `SELECT u.email_id,s.stock_name,s.stock_val FROM stockapp.user_stock us
			JOIN stockapp.users u
			ON u.user_id=us.user_id
			JOIN stockapp.stocks s
			ON us.stock_id=s.stock_id
			WHERE u.user_id=?`
	conn := sql.Connmanager
	rows, err := conn.Select(query, user_email)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	var subscribedStocks []SubscribedStock
	for rows.Next() {
		var email, stock_name, val string
		if err := rows.Scan(&email, &stock_name, &val); err != nil {
			return nil, err
		}
		stockdetail := SubscribedStock{
			EmailId:   email,
			StockName: stock_name,
			Value:     val,
		}
		subscribedStocks = append(subscribedStocks, stockdetail)
	}
	return &subscribedStocks, nil
}
