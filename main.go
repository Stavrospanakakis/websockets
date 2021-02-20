package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func handleLandpage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "public/index.html")
}

func handleWebSockets(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	conn, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Println("upgrade", err)
	}
	tick := time.Tick(time.Second * 2)

	for range tick {
		err = conn.WriteMessage(1, []byte("Message from server!"))
		if err != nil {
			log.Println("write:", err)
			conn.Close()
			break
		}
		_, message, err := conn.ReadMessage()

		if err != nil {
			log.Println("read:", err)
			conn.Close()
			break
		}

		fmt.Println(string(message))
	}

}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", handleLandpage)
	router.HandleFunc("/ws", handleWebSockets)

	log.Println("Server running in port 8080")
	log.Println(http.ListenAndServe(":8080", router))
}
