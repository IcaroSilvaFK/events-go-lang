package rabbitmq

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func OpenChannel() (*amqp.Channel){

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")

	if err != nil {
		log.Fatal(err)
	}

	ch, err := conn.Channel()

	
	if err != nil {
		log.Fatal(err)
	}


	return ch
}

func Consume(ch *amqp.Channel, out chan<- amqp.Delivery,queue string) error {

	msgs,err := ch.Consume(
		queue,
		"go-consumer",
		false,
		false,
		false,
		false,
		nil,
	)


	if err != nil {
		return err
	}

	for msg := range msgs {
		out <- msg
	}

	return nil

}

func Publish(ch *amqp.Channel, body string,queue string) bool {

	err := ch.Publish(
		queue,
		"",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body: []byte(body),
		},
	)

	return err == nil 
}
