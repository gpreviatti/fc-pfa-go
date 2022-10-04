package database

import (
	"context"
	"os"

	"testing"

	"github.com/google/uuid"
	"github.com/gpreviatti/fc-pfa-go/internal/order/entity"
	"github.com/gpreviatti/fc-pfa-go/internal/order/infra/database"
	"github.com/gpreviatti/fc-pfa-go/pkg/mongodb"
	"github.com/stretchr/testify/assert"
)

func Should_Insert_Order_In_Collection_With_Sucess(t *testing.T) {
	// Arrange
	os.Setenv("MONGO_CONNECTION_STRING", "mongodb://root:root@mongo:27017")
	ctx := context.TODO()
	client := mongodb.GetConnection(ctx)
	db := mongodb.GetDatabase(client, "pfa_go")

	repository := NewOrderRepository(db)
	order, _ := entity.NewOrder(uuid.New().String(), 10, 10)

	// Act
	result := repository.Save(ctx, order)

	// Assert
	assert.NoError(t, result)
}
