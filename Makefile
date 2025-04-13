.PHONY: help
help:
	@echo "Welcome to helper of Makefile"
	@echo "Use `make <target>` where <target> is one of:"
	@echo "key-gen	creating a couple of keys and login and password files in secret directory"


.PHONY: key-gen
key-gen: secret

secret: create-dir gen-private-key gen-public-key

create-dir:
	@mkdir -p certs
gen-private-key:
	@openssl genrsa -out certs/private_key.pem 2048
gen-public-key:
	@openssl rsa -in certs/private_key.pem -pubout -outform PEM -out certs/public_key.pem

.PHONY: lines
lines:
	git ls-files | xargs wc -l

.PHONY: build
build:
	go build -o app.o ./cmd/app/