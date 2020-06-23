package createuser

import (
	"github.com/JohnGeorge47/stock-application/internal/models"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type Signup struct {
	EmailId     string `json:"email_id"`
	UserName    string `json:"user_name"`
	Password    string  `json:"password"`
}

const createdFormat = "2006-01-02 15:04:05"


func Create(user Signup)error{
	userdata:=models.User{
		EmailId:    user.EmailId,
		UserName:   user.UserName,
	}
	lastinsert,err:=models.CreateUser(userdata)
	if err!=nil{
		return err
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	pwdData:=models.UserPassword{
		UserId:     string(*lastinsert),
		UserEmail:   user.UserName,
		Password:   string(hash),
		Updated_at: time.Unix(1391878657, 0).Format(createdFormat),
	}
	err=models.PasswordCreate(pwdData)
	if err!=nil{
		return err
	}
	return nil
}

