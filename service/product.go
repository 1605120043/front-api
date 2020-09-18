package service

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/gin-gonic/gin"
	"github.com/shinmigo/pb/productpb"
)

type Product struct {
	*gin.Context
}

func NewProduct(c *gin.Context) *Product {
	return &Product{Context: c}
}

func (m *Product) Index() (*productpb.ListProductRes, error) {
	spew.Dump("ok")
	return nil, nil
}
