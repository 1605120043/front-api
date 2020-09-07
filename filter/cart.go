package filter

import (
	"github.com/gin-gonic/gin"
	"goshop/front-api/pkg/validation"
)

type Cart struct {
	validation validation.Validation
	*gin.Context
}

func NewCart(c *gin.Context) *Cart {
	return &Cart{Context: c, validation: validation.Validation{}}
}

func (m *Cart) Add() error {
	return nil
}

func (m *Cart) Delete() error {
	return nil
}

func (m *Cart) Index() (map[string]interface{}, error) {
	return nil, nil
}
