package banner

import (
	"github.com/gin-gonic/gin"
	"supermarket/app/models"
	"supermarket/app/response"
)

/**
 * 获取商品列表
 */
func GetBanners(c *gin.Context) {
	// 申明banner
	var BannerModel models.Banner
	// 获取请求参数
	bannerType := c.Param("type")
	// 判断获取的banner类型
	banners := BannerModel.GetBannerByType(bannerType)
	// 响应
	resGin := response.Gin{C: c}

	resGin.Success("OK", banners)

}
