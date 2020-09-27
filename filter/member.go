package filter

import (
	"github.com/gin-gonic/gin"
	"github.com/shinmigo/pb/memberpb"
	"goshop/front-api/pkg/validation"
	"goshop/front-api/service"
)

type Member struct {
	validation validation.Validation
	*gin.Context
}

func NewMember(c *gin.Context) *Member {
	return &Member{Context: c, validation: validation.Validation{}}
}

func (m *Member) Info() (*memberpb.LoginRes, error) {
	return service.NewMember(m.Context).Info()
}
