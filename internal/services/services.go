package services

import (
	"encoding/json"
	"github.com/nats-io/stan.go"
	"l0/internal/database"
	"l0/internal/models"
	"log"
)

func unmarshalMessage(msg *stan.Msg) (models.Order, error) {
	var order models.Order

	err := json.Unmarshal(msg.Data, &order)
	if err != nil {
		return models.Order{}, err
	}

	return order, nil
}

// GetMessage from the channel
func GetMessage() {

	sc, _ := stan.Connect("prod", "sub1")

	defer sc.Close()

	_, err := sc.Subscribe("msg", func(m *stan.Msg) {
		order, err := unmarshalMessage(m)
		if err != nil {
			log.Printf("Error in marshaling message: %v\n", err)
		} else {
			err = database.AddMessageToDatabase(order) // add to database
			if err != nil {
				log.Print(err)
			} else {
				models.Cache[order.OrderUID] = order // add to cache
			}
		}
	})
	if err != nil {
		log.Printf("Error in subscription: %v\n", err)
	}

	// block routine for getting messages
	a := make(chan struct{})
	<-a
}
