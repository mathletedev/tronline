package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var (
	addr     = flag.String("addr", "0.0.0.0:8080", "server address")
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

func handler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("! [connect] ", err)
		return
	}
	defer conn.Close()

	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println("! [read] ", err)
			return
		}
		log.Println("[recieve] ", p)
		err = conn.WriteMessage(messageType, p)
		if err != nil {
			log.Println("! [write] ", err)
			return
		}
	}
}

func main() {
	http.HandleFunc("/", handler)

	fmt.Println("listening to", *addr)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
