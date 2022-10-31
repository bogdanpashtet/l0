go-install:
	go get github.com/nats-io/stan.go/
	go get github.com/jackc/pgx/v5
	go get github.com/gorilla/mux

run-nats-server:
	nats-streaming-server -cid prod