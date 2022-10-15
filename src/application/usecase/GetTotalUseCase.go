package usecase

import (
	"context"

	"github.com/gpreviatti/fc-pfa-go/application/dtos"
	"github.com/gpreviatti/fc-pfa-go/domain/interfaces"
)


type GetTotalUseCase struct {
	OrderRepository interfaces.OrderRepositoryInterface
}

func NewGetTotalUseCase(orderRepository interfaces.OrderRepositoryInterface) *GetTotalUseCase {
	return &GetTotalUseCase{OrderRepository: orderRepository}
}

func (c *GetTotalUseCase) Execute(ctx context.Context) (*dtos.GetTotalOutputDto, error) {
	total, err := c.OrderRepository.GetTotal(ctx)
	if err != nil {
		return nil, err
	}

	return &dtos.GetTotalOutputDto{Total: total}, nil
}
