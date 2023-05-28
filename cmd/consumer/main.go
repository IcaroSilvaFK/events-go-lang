package main

import (
	"fmt"
	"time"

	"github.com/IcaroSilvaFK/events-go-lang/pkg/rabbitmq"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main(){

	ch := rabbitmq.OpenChannel()

	defer ch.Close()

	msgs := make(chan amqp.Delivery)

	go rabbitmq.Consume(ch, msgs,"events")

	for msg := range msgs {
		bt := msg.Body

		fmt.Println(string(bt))
		msg.Ack(false)	
		time.Sleep(1 * time.Second)
	}

}