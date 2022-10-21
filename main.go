package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"site/webserver"
	"site/webserver/modules"
)

var host = "0.0.0.0"
var port = "443"
var domain = "eterinte.com"

func main() {

	ctx := context.Background()

	// чтение кофига
	conf, err := modules.ReadConfig("webserver/config/config.json")
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

	server := http.Server{
		Addr:    ":443",
		Handler: site.Router,
		TLSConfig: &tls.Config{
			NextProtos: []string{"h2", "http/1.1"},
		},
	}

	go func() {
		if err := http.ListenAndServe(":80", http.HandlerFunc(redirectTLS)); err != nil {
			log.Fatalf("ListenAndServe error: %v", err)
		}
	}()


	fmt.Printf("Server listening on %s", server.Addr)
	if err := server.ListenAndServeTLS("server.crt", "server.key"); err != nil {
		fmt.Println(err)
	}

}


func redirectTLS(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "https://"+domain+":443"+r.RequestURI, http.StatusMovedPermanently)
}
