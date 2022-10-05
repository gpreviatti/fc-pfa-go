package usecase

import (
	"context"

	"github.com/gpreviatti/fc-pfa-go/entity"
)

type GetTotalOutputDto struct {
	Total int64
}

type GetTotalUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewGetTotalUseCase(orderRepository entity.OrderRepositoryInterface) *GetTotalUseCase {
	return &GetTotalUseCase{OrderRepository: orderRepository}
}

func (c *GetTotalUseCase) Execute(ctx context.Context) (*GetTotalOutputDto, error) {
	total, err := c.OrderRepository.GetTotal(ctx)
	if err != nil {
		return nil, err
	}
	return &GetTotalOutputDto{Total: total}, nil
}
