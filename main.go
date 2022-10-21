package main

import (
	"context"
	"crypto/tls"
	"log"
	"net/http"
	"path/filepath"
	"site/webserver"
	"site/webserver/modules"

	"golang.org/x/crypto/acme/autocert"
)

var host = "0.0.0.0"
var port = "443"

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
	domain := "eterinte.com"
	// addr := fmt.Sprintf("%s:%s", host, port)
	certManager := autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist(domain),
		Cache:      autocert.DirCache("certs"),
	}

	tlsConfig := certManager.TLSConfig()
	tlsConfig.GetCertificate = getSelfSignedOrLetsEncryptCert(&certManager)
	server := http.Server{
		Addr:      ":443",
		Handler:   site.Router,
		TLSConfig: tlsConfig,
	}
	err = server.ListenAndServe()
	// log.Fatal(srv.ListenAndServeTLS("./server.crt", "./server.key"))
	// err = http.Serve(autocert.NewListener(domain), site.Router)
	// err = http.ListenAndServeTLS(addr, "./server.crt", "./server.key", site.Router)

	if err != nil {
		log.Fatal(fmt.Sprintf("Не удалось запустить сервер: %v", err))
	}

}


func getSelfSignedOrLetsEncryptCert(certManager *autocert.Manager) func(hello *tls.ClientHelloInfo) (*tls.Certificate, error) {
  return func(hello *tls.ClientHelloInfo) (*tls.Certificate, error) {
    dirCache, ok := certManager.Cache.(autocert.DirCache)
    if !ok {
      dirCache = "certs"
    }

    keyFile := filepath.Join(string(dirCache), hello.ServerName+".key")
    crtFile := filepath.Join(string(dirCache), hello.ServerName+".crt")
    certificate, err := tls.LoadX509KeyPair(crtFile, keyFile)
    if err != nil {
      fmt.Printf("%s\nFalling back to Letsencrypt\n", err)
      return certManager.GetCertificate(hello)
    }
    fmt.Println("Loaded selfsigned certificate.")
    return &certificate, err
  }
}