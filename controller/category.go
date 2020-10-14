package controller

import (
	"goshop/front-api/filter"
)

var categoryFilter *filter.Category

type Category struct {
	Base
}

func (m *Category) Initialise() {
	categoryFilter = filter.NewCategory(m.Context)
}

func (m *Category) Index() {
	str, err := categoryFilter.Index()
	if err != nil {
		m.SetResponse(nil, err)
		return
	}
	
	m.SetResponse(str)
}
