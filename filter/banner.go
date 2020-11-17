package filter

import (
	"goshop/front-api/model/banner"
	"goshop/front-api/service"

	"github.com/gin-gonic/gin"
	"github.com/shinmigo/pb/shoppb"
)

type BannerAd struct {
	*gin.Context
}

func NewBannerAd(c *gin.Context) *BannerAd {
	return &BannerAd{Context: c}
}

func (m *BannerAd) Index() (*banner.BannerAd, error) {
	req := &shoppb.ListBannerAdReq{
		Page:     1,
		PageSize: 1000,
		Id:       0,
		EleType:  0,
		Status:   shoppb.BannerAdStatus_Enabled,
	}
	return service.NewBannerAd(m.Context).Index(req)
}
