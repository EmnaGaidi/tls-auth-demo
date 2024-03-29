package main

import (
	"crypto/tls"
	"fmt"
	"net/http"
)

func main() {
	// Load server certificate and private key
	serverCert, err := tls.LoadX509KeyPair("../certs/server/server.crt", "../certs/server/server.key")
	if err != nil {
		fmt.Println("Error loading server certificates:", err)
		return
	}else{
		fmt.Println("server certificates loaded successfully")
	}

	// Create a TLS configuration
	tlsConfig := &tls.Config{
		MinVersion:   tls.VersionTLS12,
		Certificates: []tls.Certificate{serverCert},
	}

	// Create a server with TLS configuration
	server := &http.Server{
		Addr:      ":8443",
		TLSConfig: tlsConfig,
	}

	fmt.Println("Server is running on https://localhost:8443")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Received request for:", r.URL.Path)

		// Send a response back to the client
		w.Write([]byte("Hello, client! This is the server response."))
	}) 

	err = server.ListenAndServeTLS("../certs/server/server.crt", "../certs/server/server.key")
	if err != nil {
		fmt.Println("Error starting server:", err)
	}else{
		fmt.Println("server started")
	}

}
