package main

import (
	"encoding/json"
	"math/rand"
	"os"

	"github.com/google/uuid"
	"github.com/gpreviatti/fc-pfa-go/infra/rabbitmq"
	amqp "github.com/rabbitmq/amqp091-go"
)

type Order struct {
	ID    string
	Price float64
}

func GenerateOrders() Order {
	return Order{
		ID:    uuid.New().String(),
		Price: rand.Float64() * 10000,
	}
}

func main() {
	conn, err := amqp.Dial(os.Getenv("RABBITMQ_CONNECTION_STRING"))
	if err != nil {
		println(err.Error())
	}
	defer conn.Close()
	
	ch, _ := conn.Channel()
	defer ch.Close()

	for i := 0; i < 100000; i++ {
		order := GenerateOrders()
		body, err := json.Marshal(order)
		if err != nil {
			println("Error to marshal object", err.Error())
		}

		err = rabbitmq.Notify(ch, body)
		if err != nil {
			println("Error to publish message on exchange", err.Error())
		}
	}
}
