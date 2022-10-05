package api

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gpreviatti/fc-pfa-go/infra/mongodb"
	"github.com/gpreviatti/fc-pfa-go/usecase"
)

func main() {
	var serverPort = "8181"
	var ctx = context.TODO()

	http.HandleFunc("/total", func(w http.ResponseWriter, r *http.Request) {
		client := mongodb.GetConnection(ctx)
		db := mongodb.GetDatabase(client, "pfa_go")
	
		repository := mongodb.NewOrderRepository(db)
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

	println("Server running on http://localhost:" + serverPort)
	http.ListenAndServe(":" + serverPort, nil)
}
