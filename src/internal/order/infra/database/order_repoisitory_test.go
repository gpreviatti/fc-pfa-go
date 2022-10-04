package database

import (
	"context"
	"testing"

	"github.com/devfullcycle/pfa-go/internal/order/infra/database"
	"go.mongodb.org/mongo-driver/mongo"
)

func Should_Insert_Order_In_Collection() {
	// Arrange
	var ctx := context.TODO()
	var mongodb := pkg.mongodb.GetConnection(ctx)
	var repository := database.NewOrderRepository(mongodb)

	// Act
	var result = database.Save(repository, ctx)


	// Assert
}
