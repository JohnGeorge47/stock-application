package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/JohnGeorge47/stock-application/cmd/http/handlers"
	"github.com/JohnGeorge47/stock-application/internal/configmanager"
	"github.com/JohnGeorge47/stock-application/internal/models"
	"github.com/JohnGeorge47/stock-application/internal/socket"
	"github.com/JohnGeorge47/stock-application/pkg/processor"
	"github.com/JohnGeorge47/stock-application/pkg/sql"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"strings"
	"time"
)

var (
	config = flag.String("config", "./config.json", "config file path")
	port   = flag.String("port", "3003", "Host port")
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
	err := sql.InitMysqlConn()
	if err != nil {
		log.Fatal(err, "Error connecting to mysql")
	}
	ctx := context.Background()
	hub := socket.NewHub()
	go hub.Run(ctx)
	//here im running a background ticker which will do a db update every n seconds
	ticker := time.NewTicker(1000 * time.Millisecond)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				models.UpdateStocks()
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()
	//Here we have a processor which does the background job of publishing messages to the ws
	//connection
	done := make(chan struct{})
	procesticker := time.NewTicker(1000 * time.Millisecond)
	go func() {
		for {
			select {
			case <-procesticker.C:
				processor.Worker(hub)
			case <-done:
				ticker.Stop()
				return
			}
		}
	}()
	//This is the handler for the ws
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(hub, w, r, ctx)
	})
	http.Handle("/create_user", http.HandlerFunc(handlers.SignupHandler))
	http.Handle("/login", http.HandlerFunc(handlers.LoginHandler))
	http.Handle("/validate_user", http.HandlerFunc(handlers.SessionHandler))
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
	userid := r.URL.Query().Get("email_id")
	session_id := r.URL.Query().Get("session_token")
	if &userid == nil {
		fmt.Println(ctx, "No user id specified")
		ws.Close()
	}
	if &session_id == nil {
		fmt.Println(ctx, "No session_token provided specified")
		ws.Close()
	}
	ok,err:=models.CheckSession(userid,session_id)
	if err!=nil{
		fmt.Println(err)
		ws.Close()
	}
	if !*ok{
		fmt.Println("Not a valid session_id")
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
