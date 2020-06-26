package models

import (
	"fmt"
	"github.com/JohnGeorge47/stock-application/pkg/sql"
	"math/rand"
)

func GetAllStocks(limit int, offset int) error {
	query := "SELECT stock_name,stock_val,max_val,min_val FROM stocks"
	conn := sql.Connmanager
	rows, err := conn.Select(query)
	if err != nil {
		return err
	}
	fmt.Println(rows)
	return nil
}

func SubscribeToStock(email_id string, stock_name string) {

}

func UpdateStocks() {
	query := `INSERT INTO stocks(stock_id,stock_val) VALUES (1,?),(2,?),(3,?),(4,?),(5,?),(6,?),(7,?),(8,?)
             ON DUPLICATE KEY UPDATE stock_val=VALUES(stock_val)`
	conn := sql.Connmanager
	_, err := conn.Insert(query,
		rand.Intn(100),
		rand.Intn(100),
		rand.Intn(100),
		rand.Intn(100),
		rand.Intn(100),
		rand.Intn(100),
		rand.Intn(100),
		rand.Intn(100))
	if err != nil {
		fmt.Println(err)
	}
}
