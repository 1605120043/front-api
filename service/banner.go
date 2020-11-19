package service

import (
	"context"
	"encoding/json"
	"fmt"
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

	var httpInfo string
	if m.Request.TLS == nil {
		httpInfo = "http://"
	} else {
		httpInfo = "https://"
	}
	eleInfoList := make([]*banner.EleInfo, 0, 8)
	for k := range list.BannerAds {
		if list.BannerAds[k].EleInfo == "" {
			continue
		}
		errs := json.Unmarshal([]byte(list.BannerAds[k].EleInfo), &eleInfoList)
		if errs != nil {
			continue
		}
		for i := range eleInfoList {
			eleInfoList[i].ImageUrl = fmt.Sprintf("%s/image/get-image?name=%s", httpInfo+m.Request.Host, eleInfoList[i].ImageUrl)
		}

		listInfo := &banner.BannerDetail{
			Id:      list.BannerAds[k].Id,
			EleInfo: eleInfoList,
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
