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

	// Create a TLS configuration
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
	resp, err := client.Get("https://localhost:8443")
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer resp.Body.Close()

	// Read and print the response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	fmt.Println("Server Response:", string(body))
}