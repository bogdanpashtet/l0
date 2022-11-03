go-install-package:
	go get github.com/nats-io/stan.go/
	go get github.com/jackc/pgx/v5
	go get github.com/gorilla/mux

run-nats-streaming-server:
	nats-streaming-server -cid prod