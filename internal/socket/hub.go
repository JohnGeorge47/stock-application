package socket

import (
	"context"
	"fmt"
)

type Hub struct {
	//Registered users
	UserIds map[string]map[*Client]struct{}

	//Register requests from connections
	Register chan Subscription

	//UnRegister removes connection
	UnRegister chan Subscription

	//Broadcast will be broadcast to all connections with same userID
	Broadcast chan Message
}

func NewHub() *Hub {
	return &Hub{
		UserIds:    make(map[string]map[*Client]struct{}),
		Register:   make(chan Subscription),
		UnRegister: make(chan Subscription),
		Broadcast:  make(chan Message),
	}
}

//Run command runs the hub.This is run as a goroutine it picks messages,registration on
//etc. on the channels and performs actions
func (h *Hub) Run(ctx context.Context) {
	for {
		select {
		case s := <-h.Register:
			conn := h.UserIds[s.UserId]
			fmt.Println(ctx, "User has registered"+s.UserId)
			if conn == nil {
				conn = make(map[*Client]struct{})
				h.UserIds[s.UserId] = conn
			}
			h.UserIds[s.UserId][s.Conn] = struct{}{}
			fmt.Println(h.UserIds)

		case s := <-h.UnRegister:
			conn := h.UserIds[s.UserId]
			fmt.Println(ctx, "User has unregistered"+s.UserId)
			fmt.Println(conn)
			if conn != nil {
				if _, ok := conn[s.Conn]; ok {
					delete(conn, s.Conn)
					close(s.Conn.Send)
					if len(conn) == 0 {
						delete(h.UserIds, s.UserId)
					}
				}
			}
		case m := <-h.Broadcast:
			conn := h.UserIds[m.UserID]
			fmt.Println(conn)
			for c := range conn {
				select {
				case c.Send <- m.Data:
				default:
					close(c.Send)
					delete(conn, c)
					if len(conn) == 0 {
						delete(h.UserIds, m.UserID)
					}
				}
			}
		}
	}
}
