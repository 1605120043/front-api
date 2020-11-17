package controller

import "goshop/front-api/filter"

var BannerAdFilter *filter.BannerAd

type BannerAd struct {
	Base
}

func (m *BannerAd) Initialise() {
	BannerAdFilter = filter.NewBannerAd(m.Context)
}

func (m *BannerAd) Index() {
	list, err := BannerAdFilter.Index()
	if err != nil {
		m.SetResponse(list, err)
		return
	}

	m.SetResponse(list)
}
