package rabbitmq

import (
	"context"

	amqp "github.com/rabbitmq/amqp091-go"
)

func OpenChannel(rabbitmqUrl string) (*amqp.Channel, error) {
	conn, err := amqp.Dial(rabbitmqUrl)
	if err != nil {
		panic(err)
	}
	ch, err := conn.Channel()
	ch.Qos(100, 0, false)
	if err != nil {
		panic(err)
	}
	return ch, nil
}

func Consume(ch *amqp.Channel, out chan amqp.Delivery) error {
	msgs, err := ch.Consume(
		"orders",
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

func Notify(ch *amqp.Channel, body []byte) error {
	var ctx = context.TODO()
	err := ch.PublishWithContext(
		ctx,
		"orders_exchange", // exchange,
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
