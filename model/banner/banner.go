package banner

type BannerAd struct {
	Banner []*BannerDetail `json:"banner"`
	Ad     []*BannerDetail `json:"ad"`
}

type BannerDetail struct {
	Id      uint64 `json:"id"`
	EleInfo string `json:"ele_info"`
	TagName string `json:"tag_name"`
}
