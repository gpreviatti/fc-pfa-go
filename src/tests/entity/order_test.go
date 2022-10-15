package entity_test

import (
	"testing"

	"github.com/gpreviatti/fc-pfa-go/domain/entity"
	"github.com/stretchr/testify/assert"
)

func Test_GivenAnEmptyId_WhenCreateANewOrder_ThenShouldReceiveAndError(t *testing.T) {
	order := entity.Order{}
	
	assert.Error(t, order.IsValid(), "invalid id")
}

func Test_GivenAnEmptyPrice_WhenCreateANewOrder_ThenShouldReceiveAndError(t *testing.T) {
	order := entity.Order{ID: "123"}
	
	assert.Error(t, order.IsValid(), "invalid price")
}

func Test_GivenAnEmptyTax_WhenCreateANewOrder_ThenShouldReceiveAndError(t *testing.T) {
	order := entity.Order{ID: "123", Price: 10}
	
	assert.Error(t, order.IsValid(), "invalid tax")
}

func Test_GivenAValidParams_WhenCallNewOrder_ThenShould_ReceiveCreateOrderWithAllParams(t *testing.T) {
	order, err := entity.NewOrder("123", 10, 2)
	
	assert.NoError(t, err)
	assert.Equal(t, "123", order.ID)
	assert.Equal(t, 10.0, order.Price)
	assert.Equal(t, 2.0, order.Tax)
}

func Test_GivenAValidParams_WhenCallCalculateFinalPrice_ThenShouldCalculateFinalPriceAndSetItOnFinalPriceProperty(t *testing.T) {
	order, err := entity.NewOrder("123", 10, 2)
	assert.NoError(t, err)

	err = order.CalculateFinalPrice()
	assert.NoError(t, err)
	assert.Equal(t, 12.0, order.FinalPrice)
}
