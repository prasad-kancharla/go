package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/http2"
)

func main() {
	// command to generate a self-signed ssl certificate
	//openssl req -x509 -newkey rsa:2048 -nodes -keyout key.pem -out cert.pem -days 365

	http.HandleFunc("/students", func(w http.ResponseWriter, r *http.Request) {
		logServerDetails(r)
		fmt.Fprintf(w, "Handling students")
	})

	http.HandleFunc("/companies", func(w http.ResponseWriter, r *http.Request) {
		logServerDetails(r)
		fmt.Fprintf(w, "Handling companies")
	})

	cert := "cert.pem"
	key := "key.pem"

	tlsConfig := &tls.Config{
		MinVersion: tls.VersionTLS12,
	}

	server := &http.Server{
		Addr:      "localhost:3001",
		Handler:   nil,
		TLSConfig: tlsConfig,
	}

	http2.ConfigureServer(server, &http2.Server{})

	fmt.Println("Starting the server at port 3001")
	err := server.ListenAndServeTLS(cert, key)
	if err != nil {
		log.Fatalln("Got error:", err)
	}
}

func logServerDetails(r *http.Request) {
	fmt.Println("Http versions is:", r.Proto)
	fmt.Println(r.TLS.Version)
}
