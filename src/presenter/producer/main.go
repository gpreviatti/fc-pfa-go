package producer

import (
	"encoding/json"
	"math/rand"
	"os"

	"github.com/google/uuid"
	amqp "github.com/rabbitmq/amqp091-go"
)

type Order struct {
	ID    string
	Price float64
}

func GenerateOrders() Order {
	return Order{
		ID:    uuid.New().String(),
		Price: rand.Float64() * 1000,
	}
}

func Notify(ch *amqp.Channel, order Order) error {
	body, _ := json.Marshal(order)

	err := ch.Publish(
		"amq.direct", // exchange,
		"",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
	return err
}

func main() {
	conn, err := amqp.Dial(os.Getenv("RABBITMQ_CONNECTION_STRING"))
	if err != nil {
		println(err.Error())
	}
	defer conn.Close()
	
	ch, _ := conn.Channel()
	defer ch.Close()

	for i := 0; i < 100; i++ {
		order := GenerateOrders()
		err := Notify(ch, order)
		if err != nil {
			println(err.Error())
		}
	}
}
