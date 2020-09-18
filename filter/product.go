package filter

import (
	"github.com/gin-gonic/gin"
	"github.com/shinmigo/pb/productpb"
	"goshop/front-api/pkg/validation"
	"goshop/front-api/service"
)

type Product struct {
	validation validation.Validation
	*gin.Context
}

func NewProduct(c *gin.Context) *Product {
	return &Product{Context: c, validation: validation.Validation{}}
}

func (m *Product) Index() (*productpb.ListProductRes, error) {
	return service.NewProduct(m.Context).Index()
}
