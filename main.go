package main

import (
	"net/http"
	"log"
	"fmt"
	"context"
	"site/webserver"
	"site/webserver/modules"
)

var host = "0.0.0.0"
var port = "8080"

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

	addr := fmt.Sprintf("%s:%s", host, port)
	
	err = http.ListenAndServe(addr, site.Router)
	if err != nil {
		log.Fatal(fmt.Sprintf("Не удалось запустить сервер: %v", err))
	}


}