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

	// addr := fmt.Sprintf("%s:%s", host, port)
	

    cfg := &tls.Config{
        MinVersion:               tls.VersionTLS12,
        CurvePreferences:         []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
        PreferServerCipherSuites: true,
        CipherSuites: []uint16{
            tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
            tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
            tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
            tls.TLS_RSA_WITH_AES_256_CBC_SHA,
        },
    }
    srv := &http.Server{
        Addr:         "0.0.0.0:443",
        Handler:       site.Router,
        TLSConfig:    cfg,
        TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler), 0),
    }
    log.Fatal(srv.ListenAndServeTLS("../server.crt", "../server.key"))
	// err = http.ListenAndServe(addr, site.Router)
	// err = http.ListenAndServeTLS(addr, "../server.crt", "../server.key", site.Router)
	// if err != nil {
	// 	log.Fatal(fmt.Sprintf("Не удалось запустить сервер: %v", err))
	// }


}