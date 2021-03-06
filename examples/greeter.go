package main

import (
	"github.com/the5heepdev/eventsource"
	"log"
	"net/http"
	"time"
)

func main() {
	es := eventsource.New(nil, nil, func(c *eventsource.Consumer) {
		c.Message(eventsource.Msg("connect function example", "", ""))
	})
	defer es.Close()
	http.Handle("/", http.FileServer(http.Dir("./public")))
	http.Handle("/events", es)
	go func() {
		for {
			es.SendEventMessage("hello", "", "")
			log.Printf("Hello has been sent (consumers: %d)", es.ConsumersCount())
			time.Sleep(2 * time.Second)
		}
	}()
	log.Print("Open URL http://localhost:8080/ in your browser.")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
