package sql

import (
	"database/sql"
	"fmt"
	"github.com/JohnGeorge47/stock-application/internal/configmanager"
	_ "github.com/go-sql-driver/mysql"
)

var Connmanager Isql

func InitMysqlConn()error{
	db, err := sql.Open("mysql", Connectionstring())
	if err != nil {
		return err
	}
	db.SetMaxOpenConns(5)
	db.SetMaxIdleConns(3)
	err = db.Ping()
	if err != nil {
		return err
	}
	Connmanager = Client{DB: db}
}

func Connectionstring() string {
	c := configmanager.GetConfig()
	return fmt.Sprintf("%s:%s@/%s", c.SqlConf.DbUser, c.SqlConf.DbPwd, c.SqlConf.DbName)
}
