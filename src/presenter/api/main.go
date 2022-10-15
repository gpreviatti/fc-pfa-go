package main

import (
	"context"
	"encoding/json"
	"net/http"
	"os"

	"github.com/gpreviatti/fc-pfa-go/application/usecase"
	"github.com/gpreviatti/fc-pfa-go/infra/mongodb"
)

func main() {
	var ctx = context.TODO()

	http.HandleFunc("/total", func(w http.ResponseWriter, r *http.Request) {
		client := mongodb.GetConnection(ctx)
		database := mongodb.GetDatabase(client, "pfa_go")
	
		repository := mongodb.NewOrderRepository(database)
		uc := usecase.NewGetTotalUseCase(repository)

		output, err := uc.Execute(ctx)
		
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

	println("Server running on http://localhost:" + os.Getenv("SERVER_PORT"))
	http.ListenAndServe(":" + os.Getenv("SERVER_PORT"), nil)
}
