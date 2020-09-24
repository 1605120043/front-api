package filter

import (
	"github.com/gin-gonic/gin"
	"github.com/shinmigo/pb/shoppb"
	"goshop/front-api/pkg/validation"
	"goshop/front-api/service"
)

type Common struct {
	validation validation.Validation
	*gin.Context
}

func NewCommon(c *gin.Context) *Common {
	return &Common{Context: c, validation: validation.Validation{}}
}

func (m *Common) GetAreaList() (*shoppb.ListAreaRes, error) {
	return service.NewCommon(m.Context).GetAreaList()
}
