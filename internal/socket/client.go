package socket

import (
	"fmt"
	"github.com/JohnGeorge47/stock-application/pkg/uuid"
	"github.com/gorilla/websocket"
	"time"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10
)

//Client struct holds the websocket connection and the channel to send data
type Client struct {
	HubObj *Hub
	WS     *websocket.Conn
	Send   chan []byte
}

//Subscription holds a client object and the userid for that particular subscription
type Subscription struct {
	Conn   *Client
	UserId string
}

//WritePump sends messages from the hub to websocket connection
func (s *Subscription) WritePump() {
	ticker := time.NewTicker(pingPeriod)
	conn := s.Conn
	for {
		select {
		case mess, ok := <-conn.Send:
			if !ok {
				conn.WS.WriteMessage(websocket.CloseMessage, []byte{})
			}
			write, err := conn.WS.NextWriter(websocket.TextMessage)
			if err != nil {
				fmt.Println(err)
				return
			}
			write.Write(mess)
			err = write.Close()
			if err != nil {
				fmt.Println(err)
				return
			}
		case <-ticker.C:
			if err := conn.WS.WriteMessage(websocket.PingMessage, nil); err != nil {
				fmt.Println(err)
				return
			}
		}
	}
}

func (s Subscription) ReadPump() {
	conn := s.Conn
	defer func() {
		conn.HubObj.UnRegister <- s
	}()
	conn.WS.SetPongHandler(func(string) error {
		err := conn.WS.SetReadDeadline(time.Now().Add(pongWait))
		if err != nil {
			fmt.Println(err)
			return err
		}
		return nil
	})
	for {
		_, message, err := conn.WS.ReadMessage()
		fmt.Println(string(message))
		if err != nil {
			fmt.Println(err)
			break
		}
		m := Message{
			MessageID: uuid.GetUUID(),
			UserID:    s.UserId,
			Data:      message,
		}
		conn.HubObj.Broadcast <- m
	}
}
