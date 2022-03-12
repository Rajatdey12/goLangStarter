package main

import (
	"fmt"
	"net/http"

	"github.com/Rajatdey12/goLangStarter/pkg/handlers"
)

// port for application start up..
const port = ":8080"

// main is the application entrypoint..
func main() {
	http.HandleFunc("/ws", handlers.Ws)
	http.HandleFunc("/about", handlers.About)

	done := make(chan bool)
	go http.ListenAndServe(port, nil)
	fmt.Printf("Server started at port %v", port)
	<-done
}
