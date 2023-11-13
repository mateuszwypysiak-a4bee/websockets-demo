package websocket

import "fmt"

type Pool struct {
	Register   chan *Client
	Unregister chan *Client
	Clients    map[*Client]bool
	Broadcast  chan Message
}

func NewPool() *Pool {
	return &Pool{
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[*Client]bool),
		Broadcast:  make(chan Message),
	}
}

func (pool *Pool) Start() {
	for {
		select {
		case client := <-pool.Register:
			pool.Clients[client] = true
			fmt.Printf("New client {%s} registered. Size of Connection Pool: %d", client.Name, len(pool.Clients))
			for c := range pool.Clients {
				fmt.Println(c)
				c.Connection.WriteJSON(Message{Type: 1, Body: client.Name + " joined"})
			}
		case client := <-pool.Unregister:
			delete(pool.Clients, client)
			fmt.Printf("Client {%s} unregistered. Size of Connection Pool: %d", client.Name, len(pool.Clients))
			for c := range pool.Clients {
				fmt.Println(c)
				c.Connection.WriteJSON(Message{Type: 1, Body: client.Name + " left"})
			}
		case message := <-pool.Broadcast:
			fmt.Println("Sendign message to all clients in Connection Pool")
			for c := range pool.Clients {
				if err := c.Connection.WriteJSON(message); err != nil {
					fmt.Println(err)
					return
				}
			}
		}
	}
}
