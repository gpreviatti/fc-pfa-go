package usecase_test

import (
	"context"
	"testing"

	"github.com/gpreviatti/fc-pfa-go/application/dtos"
	"github.com/gpreviatti/fc-pfa-go/application/usecase"
	"github.com/gpreviatti/fc-pfa-go/domain/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type OrderRepositoryMock struct {
	mock.Mock
}

func (m *OrderRepositoryMock) Save(ctx context.Context, order *entity.Order) error {
	args := m.Called(ctx, order);
	return args.Error(0)
}

func (m *OrderRepositoryMock) GetTotal(ctx context.Context) (int64, error) {
	args := m.Called(ctx);
	return 0, args.Error(1)
}

func Test_Should_Calculate_Final_Price_With_Success(t *testing.T) {
	//Arrange
	mockRepository := new(OrderRepositoryMock)
	mockRepository.On("Save", mock.Anything).Return(nil)

	ctx := context.TODO()

	input := dtos.OrderInputDTO {
		ID:    	"1",
		Price: 	10,
		Tax: 	20,
	}
	uc := usecase.NewCalculateFinalPriceUseCase(mockRepository)

	// Act
	result, err := uc.Execute(input, ctx)

	// Assert
	assert.Nil(t, err)
	assert.NotEqual(t, "", result.ID, "result should not be nil")
}
