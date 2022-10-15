package interfaces

import (
	"context"

	"github.com/gpreviatti/fc-pfa-go/domain/entity"
)

type OrderRepositoryInterface interface {
	Save(ctx context.Context, order *entity.Order) error
	GetTotal(ctx context.Context) (int64, error)
}
