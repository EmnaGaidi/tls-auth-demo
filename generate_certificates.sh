#!/bin/bash

mkdir -p certs/ca certs/server certs/client

cd certs

# Create CA
openssl req -new -x509 -nodes -sha256 -days 365 -keyout ca/ca.key -out ca/ca.crt -subj "//C=TN//ST=Tunisia//L=Tunis//O=INSAT//CN=TLS-Demo-CA"

# Generate server certificate
openssl genrsa -out server/server.key 2048
openssl req -new -out server/server.csr -key server/server.key
openssl x509 -req -in server/server.csr -CA ca/ca.crt -CAkey ca/ca.key -CAcreateserial -out server/server.crt -days 365

# Generate client certificate
openssl genrsa -out client/client.key 2048
openssl req -new -out client/client.csr -key client/client.key
openssl x509 -req -in client/client.csr -CA ca/ca.crt -CAkey ca/ca.key -CAcreateserial -out client/client.crt -days 365
