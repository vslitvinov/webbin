package main

import (
	"net/http"
	"log"
	"fmt"
	"site/webserver"
)

var host = "0.0.0.0"
var port = "8080"

func main() {

	site := webserver.NewWebServer()

	addr := fmt.Sprintf("%s:%s", host, port)
	
	err := http.ListenAndServe(addr, site.Router)
	if err != nil {
		log.Fatal(fmt.Sprintf("Не удалось запустить сервер: %v", err))
	}


}