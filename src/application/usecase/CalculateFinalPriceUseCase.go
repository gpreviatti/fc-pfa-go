package usecase

import (
	"context"

	"github.com/gpreviatti/fc-pfa-go/application/dtos"
	"github.com/gpreviatti/fc-pfa-go/domain/entity"
	"github.com/gpreviatti/fc-pfa-go/domain/interfaces"
)

type CalculateFinalPriceUseCase struct {
	OrderRepository interfaces.OrderRepositoryInterface
}

func NewCalculateFinalPriceUseCase(orderRepository interfaces.OrderRepositoryInterface) *CalculateFinalPriceUseCase {
	return &CalculateFinalPriceUseCase{OrderRepository: orderRepository}
}

func (c *CalculateFinalPriceUseCase) Execute(input dtos.OrderInputDTO, ctx context.Context) (*dtos.OrderOutputDTO, error) {
	order, err := entity.NewOrder(input.ID, input.Price, input.Tax)
	if err != nil {
		return nil, err
	}

	err = order.CalculateFinalPrice()
	if err != nil {
		return nil, err
	}

	err = c.OrderRepository.Save(ctx, order)
	if err != nil {
		return nil, err
	}

	return &dtos.OrderOutputDTO{
		ID:         order.ID,
		Price:      order.Price,
		Tax:        order.Tax,
		FinalPrice: order.FinalPrice,
	}, nil
}
