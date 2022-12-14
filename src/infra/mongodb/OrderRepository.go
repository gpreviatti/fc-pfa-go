package mongodb

import (
	"context"

	"github.com/gpreviatti/fc-pfa-go/domain/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type OrderRepository struct {
	database *mongo.Database
}

func NewOrderRepository(database *mongo.Database) *OrderRepository {
	return &OrderRepository{database: database}
}

func (repository *OrderRepository) Save(ctx context.Context, order *entity.Order) error {
	_, error := repository.database.Collection("orders").InsertOne(ctx, order)

	return error
}

func (repository *OrderRepository) GetTotal(ctx context.Context) (int64, error) {
	total, err := repository.database.Collection("orders").CountDocuments(ctx, bson.D{{}})
	if err != nil {
		return 0, err
	}
	return total, nil
}
