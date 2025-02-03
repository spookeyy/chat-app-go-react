package main

import (
	"fmt"
	"net/http"
	"log"
	"github.com/gorilla/websocket"
)

func setupRoutes() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "Simple Server running on port 8080")
	})

	// mapping our '/ws' endpoint to the serveWs function
	http.HandleFunc("/ws", serveWs)
}


func main() {
	setupRoutes()
	http.ListenAndServe(":8080", nil)
}

// explanation
// fmt is a package that provides functions for formatting data
// Fprintf formats according to a format specifier and writes to w (which is usually os.Stdout)

// http is a package that provides functions for handling HTTP requests and responses
// ListenAndServe starts a server and handles incoming requests on the specified address
// nil is a pointer to a struct that has no fields and is used to represent a nil value in Go

// fmt.Fprintf(w, "Simple Server") writes the string "Simple Server" to the response writer w using the Fprintf function
// r is a pointer to a struct that contains information about the request such as the request method, URL, and headers


// defining an upgrader
// an upgrader is a struct that provides a way to upgrade an HTTP connection to a WebSocket connection 
// it is used to upgrade the HTTP connection to a WebSocket connection
// the Upgrade method is used to upgrade the HTTP connection to a WebSocket connection

var upgrader = websocket.Upgrader {
	ReadBufferSize: 1024, // this is the size of the buffer used to store incoming messages
	WriteBufferSize: 1024, // this is the size of the buffer used to store outgoing messages

	CheckOrigin: func(r *http.Request) bool { return true}, // this is a function that is called to check if the request is from a trusted origin
}

// defining a reader that will listen for 
// new messages being ent to our websocket endpoint

func reader( conn *websocket.Conn ) {
	for {
		// read in message
		messageType, p, err := conn.ReadMessage()

		if err != nil {
			log.Println(err)
			return
		}

		// print the message for clarity
		fmt.Println(string(p))

		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return
		}
	}

}

//defining a websocket endpoint
func serveWs(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Host)

	// upgrade the connection to a websocket connection
	// connection
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		// return
	}

	// listen indefinitely for new messages coming through
	// on our websocket connection
	reader(ws)
}

// mapping our '/ws' endpoint to the serverWs function

