package main

import (
	"context"
	"encoding/json"
	"net/http"
	"os"

	"github.com/gpreviatti/fc-pfa-go/internal/order/infra/database"
	"github.com/gpreviatti/fc-pfa-go/internal/order/usecase"
	"github.com/gpreviatti/fc-pfa-go/pkg/mongodb"
)

func main() {
	var serverPort = "8181"
	var ctx = context.TODO()
	var mysqlUrl = "root:root@tcp(mysql:3306)/orders"
	if os.Getenv("ENVIRONMENT") == "Production" {
		mysqlUrl = "root:root@tcp(mysql-service.pfa-go.svc:3306)/orders"
	}

	http.HandleFunc("/total", func(w http.ResponseWriter, r *http.Request) {
		client := mongodb.GetConnection(ctx)
		db := mongodb.GetDatabase(client, "pfa_go")
	
		repository := database.NewOrderRepository(db)
		uc := usecase.NewGetTotalUseCase(repository)

		output, err := uc.Execute()
		
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			println(err.Error())
			return
		}

		json.NewEncoder(w).Encode(output)
	})

	http.HandleFunc("/hc", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode("Healthy")
	})

	println("Server running on http://localhost:" + serverPort)
	http.ListenAndServe(":" + serverPort, nil)
}
