package database

import (
	"context"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gpreviatti/fc-pfa-go/internal/order/entity"
	"go.mongodb.org/mongo-driver/mongo"
)

type OrderRepository struct {
	database *mongo.Database
}

func NewOrderRepository(database *mongo.Database) *OrderRepository {
	return &OrderRepository{database: database}
}

func (repository *OrderRepository) Save(ctx context.Context, order *entity.Order) error {
	_, err := repository.database.Collection("orders").InsertOne(ctx, order)

	return err
}

// func (r *OrderRepository) GetTotal() (int, error) {
// 	var total int
// 	err := r.Db.QueryRow("SELECT COUNT(*) FROM orders").Scan(&total)
// 	if err != nil {
// 		return 0, err
// 	}
// 	return total, nil
// }
