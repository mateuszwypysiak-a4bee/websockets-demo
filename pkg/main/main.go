package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"a4bee.com/websocket/pkg/static"
	"a4bee.com/websocket/pkg/websocket"
)

func main() {
	fmt.Println("Websockets backend")
	router := mux.NewRouter()

	router.PathPrefix("/static/").Handler(static.HandleStaticFiles())
	// router.PathPrefix("/native").HandlerFunc(static.HandleNative())
	router.HandleFunc("/ws", websocket.HandleWS())

	http.ListenAndServe("localhost:8080", router)
}
