package mongodb

import (
	"context"
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/gpreviatti/fc-pfa-go/entity"
	"github.com/stretchr/testify/assert"
)

func Test_Should_Insert_Order_In_Collection_With_Sucess(t *testing.T) {
	// Arrange
	os.Setenv("MONGO_CONNECTION_STRING", "mongodb://root:root@mongo:27017")
	ctx := context.TODO()
	client := GetConnection(ctx)
	db := GetDatabase(client, "pfa_go")

	repository := NewOrderRepository(db)
	order, _ := entity.NewOrder(uuid.New().String(), 10, 10)

	// Act
	result := repository.Save(ctx, order)

	// Assert
	assert.NoError(t, result)
}
