package crawler

import (
	"encoding/json"
	"github.com/streadway/amqp"
	"log"
)

type Crawler struct {
	Fetcher      Fetcher
	MessageChan  *amqp.Channel
	MessageQueue *amqp.Queue
}

func (c *Crawler) Publish(message interface{})  {
	messageJson, err := json.Marshal(message)
	if err != nil {
		log.Fatal()
	}
	err = c.MessageChan.Publish(
		"",     // exchange
		c.MessageQueue.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing {
			ContentType: "text/plain",
			Body:        messageJson,
		})
}

func (c *Crawler) Start() error {
	ch := make(chan interface{})
	go c.Fetcher.Fetch(ch)
	log.Println("Fetching started")

	for message := range ch {
		c.Publish(message)
	}
	return nil
}