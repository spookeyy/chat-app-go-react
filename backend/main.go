package main

import (
	// "fmt"
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

func setupRoutes(router *gin.Engine) {
	router.GET("/", func(c *gin.Context) {
		c.String(200, "Simple Server running on port 8080")
	})

	// mapping our '/ws' endpoint to the serveWs function
	router.GET("/ws", func(c *gin.Context) {
		serveWs(c)
	})
}

func main() {
	router := gin.Default()
	setupRoutes(router)
	router.Run(":8080") // This replaces http.ListenAndServe
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

//defining a websocket endpoint
// Gin-style WebSocket handler
func serveWs(c *gin.Context) {
	// Upgrade HTTP connection to WebSocket
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Printf("WebSocket upgrade error: %v", err)
		c.AbortWithStatusJSON(500, gin.H{"error": "WebSocket upgrade failed"})
		return
	}
	defer ws.Close()

	// Start reading messages
	for {
		messageType, message, err := ws.ReadMessage()
		if err != nil {
			log.Printf("WebSocket read error: %v", err)
			break
		}

		// Log received message
		log.Printf("Received: %s", string(message))

		// Echo the message back
		if err := ws.WriteMessage(messageType, message); err != nil {
			log.Printf("WebSocket write error: %v", err)
			break
		}
	}
}

// mapping our '/ws' endpoint to the serverWs function