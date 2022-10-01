package main

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/devfullcycle/pfa-go/internal/order/infra/database"
	"github.com/devfullcycle/pfa-go/internal/order/usecase"
)

func main() {
	http.HandleFunc("/total", func(w http.ResponseWriter, r *http.Request) {
		db, _ := sql.Open("mysql", "root:root@tcp(mysql:3306)/orders")
		defer db.Close()
	
		repository := database.NewOrderRepository(db)
		uc := usecase.NewGetTotalUseCase(repository)

		output, err := uc.Execute()
		
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(output)
	})

	http.HandleFunc("/hc", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode("Healthy")
	})

	go http.ListenAndServe(":8181", nil)
}