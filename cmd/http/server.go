package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/JohnGeorge47/stock-application/internal/configmanager"
	"github.com/JohnGeorge47/stock-application/internal/socket"
	"github.com/JohnGeorge47/stock-application/pkg/redis"
	"github.com/gorilla/websocket"
	"github.com/JohnGeorge47/stock-application/cmd/http/handlers"
	"log"
	"net/http"
	"strings"
	"github.com/JohnGeorge47/stock-application/pkg/sql"
)

var (
	config = flag.String("config", "./config.json", "config file path")
	port   = flag.String("port", "8000", "Host port")
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		for _, origin := range configmanager.GetConfig().Origins {
			host := r.Header.Get("Origin")
			fmt.Println(host)
			if strings.Contains(host, origin) {
				return true
			}
		}
		return true
	},
}

func main() {
	if err := configmanager.InitConfig(config); err != nil {
		fmt.Println(err)
	}
	//redis.InitConnection()
	err:=sql.InitMysqlConn()
	if err!=nil{
		log.Fatal(err,"Error connecting to mysql")
	}
	ctx := context.Background()
	hub := socket.NewHub()
	go hub.Run(ctx)
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(hub, w, r, ctx)
	})
	http.Handle("/create_user",http.HandlerFunc(handlers.SignupHandler))
	fmt.Println("server Listening on", *port)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", *port), nil); err != nil {
		fmt.Println(err)
	}
}

func serveWs(hub *socket.Hub, w http.ResponseWriter, r *http.Request, ctx context.Context) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
	}
	userid := r.URL.Query().Get("user_id")
	session_id := r.URL.Query().Get("session_id")
	if &userid == nil {
		fmt.Println(ctx, "No user id specified")
		ws.Close()
	}
	if &session_id == nil {
		fmt.Println(ctx, "No user id specified")
		ws.Close()
	}
	conn := &socket.Client{
		WS:     ws,
		Send:   make(chan []byte, 256),
		HubObj: hub,
	}

	sub := socket.Subscription{
		Conn:   conn,
		UserId: userid,
	}
	if &userid != nil {
		hub.Register <- sub
		go sub.WritePump()
		go sub.ReadPump()
	}
}
