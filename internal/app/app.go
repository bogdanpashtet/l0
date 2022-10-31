package app

import (
	"github.com/gorilla/mux"
	"l0/internal/database"
	"l0/internal/services"
	"l0/internal/transport"
	"log"
	"net/http"
)

func Run() {
	err := database.SyncCacheAndDatabase()
	if err != nil {
		log.Print(err)
	}

	router := mux.NewRouter()
	router.HandleFunc("/", transport.MainPage)
	router.HandleFunc("/message/{uid:[0-9]+}", transport.MessageHandler)

	go services.GetMessage()

	err = http.ListenAndServe(transport.Port, router)
	if err != nil {
		log.Printf("Error in lauching server: %v", err)
	}
}
