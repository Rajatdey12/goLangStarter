package main

import (
	"fmt"
	"net/http"

	"github.com/Rajatdey12/goLangStarter/pkg/config"
	"github.com/Rajatdey12/goLangStarter/pkg/utils"
)

// port for application start up..
const port = ":8080"

// Setting up the appConfig
var appConfig config.AppConfig

// main is the application entrypoint..
func main() {

	appConfig.InProduction = false
	appConfig.Session = utils.CreateSession(appConfig)

	http.Handle("/", routes())

	done := make(chan bool)
	go http.ListenAndServe(port, nil)
	fmt.Println("Server started at port ==>", port)
	<-done
}
