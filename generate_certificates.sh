#!/bin/bash

mkdir -p certs/ca certs/server certs/client

cd certs

# Create CA
openssl req -new -x509 -nodes -sha256 -days 365 -extensions v3_ca -keyout ca/ca.key -out ca/ca.crt -subj "//C=TN//ST=Tunisia//L=Tunis//O=INSAT//CN=TLS-Demo-CA"

# Generate server certificate
openssl genrsa -out server/server.key 2048
openssl req -new -out server/server.csr -key server/server.key -config openssl-server.conf -sha256
openssl x509 -req -in server/server.csr -CA ca/ca.crt -CAkey ca/ca.key -CAcreateserial -out server/server.crt -days 365 -extensions req_ext -extfile openssl-server.conf -sha256

# Generate client certificate
openssl genrsa -out client/client.key 2048
openssl req -new -out client/client.csr -key client/client.key -config openssl-client.conf -sha256
openssl x509 -req -in client/client.csr -CA ca/ca.crt -CAkey ca/ca.key -CAcreateserial -out client/client.crt -days 365 -extensions req_ext -extfile openssl-client.conf -sha256
