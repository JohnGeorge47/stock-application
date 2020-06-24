package processor

import (
	"encoding/json"
	"fmt"
	"github.com/JohnGeorge47/stock-application/internal/models"
	"github.com/JohnGeorge47/stock-application/internal/socket"
	"github.com/JohnGeorge47/stock-application/pkg/uuid"
)

type SocketMessage struct {
	Message []models.SubscribedStock  `json:"message"`
}

func Worker(hub *socket.Hub) {
	user_list, err := models.GetLoggedinUsers()
	fmt.Println(user_list)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(user_list)
	for _, user := range *user_list {
		result, err := models.GetAllUserSubscribedStocks(user)
		if err!=nil{
			fmt.Println(err)
		}
		fmt.Println(result, err)
		if len(*result) != 0 {
			val:=*result
			messageToMarshal:=SocketMessage{Message:val}
			jsonmessage,err:=json.Marshal(messageToMarshal)
			usermail:=val[0].EmailId
			if err!=nil{
				fmt.Println(err)
			}
			mess:=socket.Message{
				MessageID: uuid.GetUUID(),
				UserID:    usermail,
				Data:      jsonmessage,
			}
			hub.Broadcast<-mess
		}
	}
}
