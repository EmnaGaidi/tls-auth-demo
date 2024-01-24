package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// Load CA certificate
    caCert, err := ioutil.ReadFile("../certs/ca/ca.crt")
    if err != nil {
		fmt.Println("Error reading CA certificate:", err)
    }

	// Create certificate pool and add caCert to it; it will be used for server certificate verification
	// Certificate pool is a collection of trusted certificates, typically used for verifying the authenticity of the server's certificate during a connection establishment
	caCertPool := x509.NewCertPool()
    caCertPool.AppendCertsFromPEM(caCert)

	// Load client certificate and private key
	clientCert, err := tls.LoadX509KeyPair("../certs/client/client.crt", "../certs/client/client.key")
	if err != nil {
		fmt.Println("Error loading client certificates:", err)
		return
	}else{
		fmt.Println("client certificates loaded successfully")
	}

	// Create a TLS configuration in order to set up a secure connection
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{clientCert},
		RootCAs:      caCertPool,
	}

	// Create an HTTP client with the TLS configuration
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: tlsConfig,
		},
	}

	// Make a request to the server
	// CN (common name) and SAN (SubjectAltName) in the server cetificate should be both = localhost
	resp, err := client.Get("https://localhost:8443/")
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}else{
		fmt.Println("Successful request : server verified by the CA")
	}
	// Defers the closing of the response body until the surrounding function (main) exits
	defer resp.Body.Close()

	// Read and print the response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	fmt.Println("Server Response:", string(body))
}