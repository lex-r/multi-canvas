package main

import (
	"net/http"
	"github.com/gorilla/websocket"
	"log"
	"github.com/lex-r/multi-canvas/messages"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func handler(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")

	if len(id) == 0 {
		log.Printf("Connection without id")
		w.WriteHeader(400)
		return
	}

	log.Printf("Handle request with id %v", id)

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	c := &connection{send: make(chan []byte, 256), ws: conn}
	monitor := &Monitor{
		WorldId: id,
		Conn: c,
	}
	monitors[c] = monitor
	h.register <- c
	go c.writePump()

	log.Printf("New connection");
	log.Printf("Count connections: %s", len(h.connections));

	c.readPump()
}

var service = Service{functions:make(map[string]func(*connection, *messages.ServerRequest))}
var monitors = map[*connection]*Monitor{}
var worlds = map[string]*World{}

func main() {

	monitors = make(map[*connection]*Monitor)
    worlds = make(map[string]*World)

	registerFunctions(&service)

	go h.run()
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/ws", handler)

	err := http.ListenAndServe(":8090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
