package main

import (
	"database/sql"

	"github.com/devfullcycle/pfa-go/internal/order/infra/database"
	"github.com/devfullcycle/pfa-go/internal/order/usecase"
	"github.com/google/uuid"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(mysql:3306)/orders")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	repository := database.NewOrderRepository(db)
	uc := usecase.NewCalculateFinalPriceUseCase(repository)
	input := usecase.OrderInputDTO{
		ID:    uuid.New().String(),
		Price: 100,
		Tax:   10,
	}
	output, err := uc.Execute(input)
	if err != nil {
		panic(err)
	}
	println(output.FinalPrice)
}
