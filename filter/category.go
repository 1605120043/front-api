package filter

import (
	"github.com/gin-gonic/gin"
	"github.com/shinmigo/pb/productpb"
	"goshop/front-api/pkg/validation"
	"goshop/front-api/service"
)

type Category struct {
	validation validation.Validation
	*gin.Context
}

func NewCategory(c *gin.Context) *Category {
	return &Category{Context: c, validation: validation.Validation{}}
}

func (m *Category) Index() (*productpb.ListCategoryRes, error) {
	listCategoryReq := &productpb.ListCategoryReq{
		Page:     1,
		PageSize: 1000,
	}
	return service.NewCategory(m.Context).Index(listCategoryReq)
}
