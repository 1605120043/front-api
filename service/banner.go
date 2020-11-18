package service

import (
	"context"
	"goshop/front-api/model/banner"
	"goshop/front-api/pkg/grpc/gclient"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/shinmigo/pb/shoppb"
)

type BannerAd struct {
	*gin.Context
}

func NewBannerAd(c *gin.Context) *BannerAd {
	return &BannerAd{Context: c}
}

func (m *BannerAd) Index(param *shoppb.ListBannerAdReq) (bannerAdList *banner.BannerAd, err error) {
	bannerAdList = &banner.BannerAd{}
	bannerAdList.Ad = make([]*banner.BannerDetail, 0, 8)
	bannerAdList.Banner = make([]*banner.BannerDetail, 0, 8)

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	list, err := gclient.BannerAdClient.GetBannerAdList(ctx, param)
	cancel()
	if err != nil {
		return bannerAdList, err
	}
	if list.Total == 0 {
		return bannerAdList, nil
	}

	for k := range list.BannerAds {
		listInfo := &banner.BannerDetail{
			Id:      list.BannerAds[k].Id,
			EleInfo: list.BannerAds[k].EleInfo,
			TagName: list.BannerAds[k].TagName,
		}

		switch list.BannerAds[k].EleType {
		case 1:
			bannerAdList.Banner = append(bannerAdList.Banner, listInfo)
		case 2:
			bannerAdList.Ad = append(bannerAdList.Ad, listInfo)
		}
	}

	return
}
