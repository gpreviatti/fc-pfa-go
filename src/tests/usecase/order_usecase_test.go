package usecase_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/google/uuid"
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
	args := m.Called(ctx, order)
	return args.Error(0)
}

func (m *OrderRepositoryMock) GetTotal(ctx context.Context) (int64, error) {
	args := m.Called(ctx)
	return int64(args.Int(0)), args.Error(1)
}

func Test_Should_Calculate_Final_Price_With_Success(t *testing.T) {
	//Arrange
	mockRepository := new(OrderRepositoryMock)
	mockRepository.On("Save", mock.Anything, mock.Anything).Return(nil)

	ctx := context.TODO()

	input := dtos.OrderInputDTO {
		ID:    	uuid.New().String(),
		Price: 	10,
		Tax: 	20,
	}
	uc := usecase.NewCalculateFinalPriceUseCase(mockRepository)

	// Act
	result, err := uc.Execute(input, ctx)
	if err != nil {
		fmt.Println(err.Error())
	}

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, result.ID, input.ID)
}


func Test_Should_Get_Total_With_Success(t *testing.T) {
	//Arrange
	const total = 100
	mockRepository := new(OrderRepositoryMock)
	mockRepository.On("GetTotal", mock.Anything).Return(total, nil)

	ctx := context.TODO()
	uc := usecase.NewGetTotalUseCase(mockRepository)
	
	// Act
	result, err := uc.Execute(ctx)
	if err != nil {
		fmt.Println(err.Error())
	}

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, int64(total), result.Total)
}
