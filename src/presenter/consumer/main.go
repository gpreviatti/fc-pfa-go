package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/gpreviatti/fc-pfa-go/order/infra/database"
	"github.com/gpreviatti/fc-pfa-go/order/usecase"
	"github.com/gpreviatti/fc-pfa-go/pkg/mongodb"
	"github.com/gpreviatti/fc-pfa-go/pkg/rabbitmq"
	amqp "github.com/rabbitmq/amqp091-go"
)

var ctx = context.TODO()

func main() {
	
	client := mongodb.GetConnection(ctx)
	db := mongodb.GetDatabase(client, "pfa_go")
	wg := sync.WaitGroup{}
	maxWorkers := 5
	
	repository := database.NewOrderRepository(db)
	uc := usecase.NewCalculateFinalPriceUseCase(repository)

	ch, err := rabbitmq.OpenChannel(os.Getenv("RABBITMQ_CONNECTION_STRING"))
	if err != nil {
		panic(err)
	}

	defer ch.Close()
	out := make(chan amqp.Delivery)
	go rabbitmq.Consume(ch, out)
	for i := 0; i < maxWorkers; i++ {
		wg.Add(1)
		i := i
		go func() {
			fmt.Println("Starting worker", i)
			defer wg.Done()
			worker(out, uc, i)
		}()
	}
	wg.Wait()
}

func worker(deliveryMessage <-chan amqp.Delivery, uc *usecase.CalculateFinalPriceUseCase, workerId int) {
	for msg := range deliveryMessage {
		var input usecase.OrderInputDTO
		err := json.Unmarshal(msg.Body, &input)
		if err != nil {
			fmt.Println("Error unmarshalling message", err)
		}
		input.Tax = 10.0

		_, err = uc.Execute(input)
		if err != nil {
			fmt.Println("Error unmarshalling message", err)
		}

		msg.Ack(false)
		fmt.Println("Worker", workerId, "processed order", input.ID)
		time.Sleep(1 * time.Second)
	}
}
