package application_test

import (
	"github.com/polianalima/go-hexagonal/application"
	"github.com/stretchr/testify/require"

	"github.com/satori/go.uuid"

	"testing"
)

func TestProduct_Enable(t *testing.T) {
	product := application.Product{}
	product.Name = "Hello"
	product.Status = application.DISABLED
	product.Price = 10

	err := product.Enable()
	require.Nil(t, err)

	product.Price = 0
	err = product.Enable()
	require.Equal(t, "The price must be greater than zero to enable the product", err.Error())
}

func TestProduct_Disable(t *testing.T) {
	product := application.Product{}
	product.Name = "Hello"
	product.Status = application.DISABLED
	product.Price = 0

	err := product.Disable()
	require.Nil(t, err)

	product.Price = 10
	err = product.Disable()
	require.Equal(t, "The price must be zero to disable the product", err.Error())
}

func TestProduct_IsValid(t *testing.T) {
	product := application.Product{}
	product.Id = uuid.NewV4().String()
	product.Name = "Hello"
	product.Status = application.DISABLED
	product.Price = 10

	_, err := product.IsValid()
	require.Nil(t, err)

	product.Price = 10
	product.Status = "INVALID"
	_, err = product.IsValid()
	require.Equal(t, "The status must be enabled or disabled", err.Error())


	product.Status = application.ENABLED
	_, err = product.IsValid()
	require.Nil(t, err)
	
	product.Price = -10
	_, err = product.IsValid()
	require.Equal(t, "The price must be greater or equal zero", err.Error())


}