run:
	echo "Run disism oauth2.1 server..." && go run ./cmd/oauth2/main.go

init:
	go mod tidy