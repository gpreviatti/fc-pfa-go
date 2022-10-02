package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/devfullcycle/pfa-go/internal/order/infra/database"
	"github.com/devfullcycle/pfa-go/internal/order/usecase"
	"github.com/devfullcycle/pfa-go/pkg/rabbitmq"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	maxWorkers := 1

	var rabbitmqUrl = "amqp://guest:guest@rabbitmq:5672"
	var mysqlUrl = "root:root@tcp(mysql:3306)/orders"

	if os.Getenv("ENVIRONMENT") == "Production" {
		rabbitmqUrl = "amqp://guest:guest@rabbitmq-service.pfa-go.svc:5672/"
		mysqlUrl = "root:root@tcp(mysql-service.pfa-go.svc:3306)/orders"
	}

	wg := sync.WaitGroup{}
	db, err := sql.Open("mysql", mysqlUrl)
	if err != nil {
		panic(err)
	}
	
	defer db.Close()
	repository := database.NewOrderRepository(db)
	uc := usecase.NewCalculateFinalPriceUseCase(repository)

	ch, err := rabbitmq.OpenChannel(rabbitmqUrl)
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
