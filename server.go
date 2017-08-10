package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("client")))
	http.HandleFunc("/connected", connected)

	err := http.ListenAndServe(":4200", nil)
	if err != nil {
		panic(err)
	}
}

var upgrader = websocket.Upgrader{} // use default options

func connected(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()

	for {
		msgtype, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}

		log.Printf("recv: %s", message)

		err = c.WriteMessage(msgtype, message)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}
