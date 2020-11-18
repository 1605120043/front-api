package service

import (
	"context"
	"fmt"
	"goshop/front-api/model/banner"
	"goshop/front-api/pkg/grpc/gclient"
	"goshop/front-api/pkg/utils"
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
		imageUrl := fmt.Sprintf("http://10.32.5.88:%s/image/get-image?name=%s", utils.C.Webhost, list.BannerAds[k].ImageUrl)
		listInfo := &banner.BannerDetail{
			Id:          list.BannerAds[k].Id,
			ImageUrl:    imageUrl,
			RedirectUrl: list.BannerAds[k].RedirectUrl,
			Sort:        list.BannerAds[k].Sort,
			TagName:     list.BannerAds[k].TagName,
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
