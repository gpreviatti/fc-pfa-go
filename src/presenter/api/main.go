package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"os"

	"github.com/devfullcycle/pfa-go/internal/order/infra/database"
	"github.com/devfullcycle/pfa-go/internal/order/usecase"
)

func main() {
	var serverPort = "8181"
	var mysqlUrl = "root:root@tcp(mysql:3306)/orders"
	if os.Getenv("ENVIRONMENT") == "Production" {
		mysqlUrl = "root:root@tcp(mysql-service.pfa-go.svc:3306)/orders"
	}

	http.HandleFunc("/total", func(w http.ResponseWriter, r *http.Request) {
		db, _ := sql.Open("mysql", mysqlUrl)
		defer db.Close()
	
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
