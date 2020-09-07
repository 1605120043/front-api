package service

import (
	"github.com/gin-gonic/gin"
)

type Cart struct {
	*gin.Context
}

func NewCart(c *gin.Context) *Cart {
	return &Cart{Context: c}
}

func (m *Cart) Add(memberId uint64) error {
	return nil
}

func (m *Cart) Delete(memberId, cartId uint64) error {
	return nil
}

func (m *Cart) Index() (map[string]interface{}, error) {
	return nil, nil
}
