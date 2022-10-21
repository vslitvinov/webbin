package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"site/webserver"
	"site/webserver/modules"

	"golang.org/x/crypto/acme/autocert"
)

var host = "0.0.0.0"
var port = "443"

func main() {

	ctx := context.Background()
	
	// чтение кофига
	conf,err  := modules.ReadConfig("webserver/config/config.json")
	if err != nil {
		log.Fatal("Initialization failed: ", err)
	}


	connConfig := fmt.Sprintf("postgres://%v:%v@%v:%v/%v", 
		conf.Database.User,
		conf.Database.Password,
		conf.Database.Host,
		conf.Database.Port,
		conf.Database.Name,
	)

	// соединение с postgres
	db, err := modules.NewPostgres(ctx, connConfig)
	if err != nil {
		log.Fatal("Postgres initialization failed: ", err)
	}
	defer db.Close()

	site := webserver.NewWebServer(db)
	domain := "eterinte.com"
	// addr := fmt.Sprintf("%s:%s", host, port)
	


    // log.Fatal(srv.ListenAndServeTLS("./server.crt", "./server.key"))
    err = http.Serve(autocert.NewListener(domain), site.Router)
	// err = http.ListenAndServeTLS(addr, "./server.crt", "./server.key", site.Router)

	if err != nil {
		log.Fatal(fmt.Sprintf("Не удалось запустить сервер: %v", err))
	}


}