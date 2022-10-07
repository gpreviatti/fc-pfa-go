package mongodb_test

import (
	"context"
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/gpreviatti/fc-pfa-go/entity"
	"github.com/gpreviatti/fc-pfa-go/infra/mongodb"
	"github.com/stretchr/testify/assert"
)

func Test_Should_Insert_Order_With_Sucess(t *testing.T) {
	// Arrange
	os.Setenv("MONGO_CONNECTION_STRING", "mongodb://root:root@mongo:27017")
	ctx := context.TODO()
	client := mongodb.GetConnection(ctx)
	db := mongodb.GetDatabase(client, "pfa_go")

	repository := mongodb.NewOrderRepository(db)
	order, _ := entity.NewOrder(uuid.New().String(), 50, 10)

	// Act
	result := repository.Save(ctx, order)

	// Assert
	assert.Nil(t, result)
}

func Test_Should_Get_Total_With_Sucess(t *testing.T) {
	// Arrange
	os.Setenv("MONGO_CONNECTION_STRING", "mongodb://root:root@mongo:27017")
	ctx := context.TODO()
	client := mongodb.GetConnection(ctx)
	db := mongodb.GetDatabase(client, "pfa_go")

	repository := mongodb.NewOrderRepository(db)

	// Act
	result, error := repository.GetTotal(ctx)

	// Assert
	assert.Nil(t, error)
	assert.NotNil(t, result)
	assert.NotEqual(t, 0, result)
}
