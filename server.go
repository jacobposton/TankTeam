package main

import (
	//"encoding/json"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

/*

Command line Args:
--------------------
-p PORT           Set port. Default is 8080
-w                Starts the web server in addition to the websocket server

*/

var port = ":8080"

var socketManager = network.wsmgr{}

func main() {
	log.Println("Starting TankTeam Server")

	fs := http.FileServer(http.Dir("res/static"))

	http.Handle("/ws", socketManager)
	http.Handle("/", fs)

	log.Printf("Server listening on port%s", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatalf("ERROR %v", err)
	}
}
