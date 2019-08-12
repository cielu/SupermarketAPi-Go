package models

import (
	_ "github.com/go-sql-driver/mysql"
)

type Banner struct {
	BannerId        uint   `json:"bannerId" gorm:"primary_key"`
	BannerName      string `json:"bannerName" gorm:"size:255"`
	ImgUrl          string `json:"imgUrl" gorm:"size:255"`
	Target          string `json:"target" gorm:"size:255"`
	ShopId          uint   `json:"shopId"`
	Type            string `json:"type"`
	Scene           string `json:"scene"`
	BackgroundStyle string `json:"backgroundStyle"`
	BackgroundClass string `json:"backgroundClass"`
	Sort            int    `json:"sort"`
}

// 获取表名
func (Banner) TableName() string {
	return "banners"
}

// 根据类型获取banner
func (banner Banner) GetBannerByType(_type string) (ret map[string]interface{}) {
	// 初始化map
	ret = make(map[string]interface{})
	// 申明 banners
	var banners, adverts []Banner
	switch _type {
	case "mobileIndex": // 小程序首页
		// 获取首页顶部banner & 首页广告
		banner.getBannerByScene("mobileIndexBanner",1,&banners)
		banner.getBannerByScene("mobileIndexAdvert",2,&adverts)
		break
	default:

	}
	ret["banners"] = banners
	ret["adverts"] = adverts
	return
}

// 查询banner
func (banner Banner) getBannerByScene(scene string,limit uint8, banners *[]Banner) {
	db.Where("scene = ?", scene).Limit(limit).Order("sort desc").Find(&banners)
}
