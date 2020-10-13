package filter

import (
	"github.com/gin-gonic/gin"
	"github.com/shinmigo/pb/productpb"
	"goshop/front-api/model/product"
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

func (m *Product) Index() ([]*product.ProductList, error) {
	categoryId := m.Query("category_id")
	tagId := m.Query("tag_id")
	page := m.DefaultQuery("page", "1")
	pageSize := m.DefaultQuery("page_size", "20")
	if len(categoryId) > 0 {
		m.validation.Numeric(categoryId).Message("分类id不正确")
	}
	
	if len(tagId) > 0 {
		m.validation.Numeric(tagId).Message("标签id不正确")
	}
	
	m.validation.Numeric(page).Message("请求参数错误")
	m.validation.Numeric(pageSize).Message("请求参数错误")
	
	if m.validation.HasError() {
		return nil, m.validation.GetError()
	}
	
	return service.NewProduct(m.Context).Index()
}

func (m *Product) Detail() (*productpb.ProductDetail, error) {
	productIdStr := m.DefaultQuery("product_id", "0")
	m.validation.Numeric(productIdStr).Message("商品id不正确")
	
	if m.validation.HasError() {
		return nil, m.validation.GetError()
	}
	
	return service.NewProduct(m.Context).Detail()
}

func (m *Product) Tag() ([]*product.TagList, error) {
	return service.NewProduct(m.Context).Tag()
}
