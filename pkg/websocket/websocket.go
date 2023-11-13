package websocket

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func HandleWS() func(w http.ResponseWriter, r *http.Request) {
	var pool = NewPool()
	go pool.Start()

	return func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		fmt.Printf("WS endpoint hit with name=%s\n", name)

		connection, err := upgrade(w, r)
		if err != nil {
			log.Println(err)
		}

		client := &Client{
			Connection: connection,
			Pool:       pool,
			Name:       name,
		}
		pool.Register <- client
		client.Read()
	}
}

func upgrade(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	connection, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return connection, err
	}
	return connection, nil

}

func reader(connection *websocket.Conn) {
	for {
		messageType, message, err := connection.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		fmt.Printf("Message: %s\n", string(message))

		if err := connection.WriteMessage(messageType, message); err != nil {
			log.Println(err)
			return
		}
	}
}

func writer(connection *websocket.Conn) {
	for {
		fmt.Println("Sending")
		messageType, r, err := connection.NextReader()
		if err != nil {
			fmt.Println(err)
			return
		}

		w, err := connection.NextWriter(messageType)
		if err != nil {
			fmt.Println(err)
			return
		}

		if _, err := io.Copy(w, r); err != nil {
			fmt.Println(err)
			return
		}

		if err := w.Close(); err != nil {
			fmt.Println(err)
			return
		}

	}
}
