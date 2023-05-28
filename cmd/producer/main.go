package main

import (
	"github.com/IcaroSilvaFK/events-go-lang/pkg/rabbitmq"
)

func main(){

	ch := rabbitmq.OpenChannel()

	defer ch.Close()

	for i:= 0; i < 1_000; i++ {
		rabbitmq.Publish(ch, "hello world","amq.direct")
	}

}