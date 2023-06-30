
all:
	openssl genrsa -out ca.key 2048
	openssl req -new -key ca.key -out ca.csr -config openssl_san.cnf
	openssl x509 -req -in ca.csr -out ca.crt -signkey ca.key -extensions v3_req -extfile openssl_san.cnf -days 3650

	openssl genrsa -out server.key 2048
	openssl req -new -key server.key -out server.csr -config openssl_san.cnf
	openssl x509 -req -sha256 -CA ca.crt -CAkey ca.key -CAcreateserial -in server.csr -out server.crt -extensions v3_req -extfile openssl_san.cnf -days 3650

	rm *.csr