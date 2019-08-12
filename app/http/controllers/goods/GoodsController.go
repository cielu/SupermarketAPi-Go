package goods

import (
	"github.com/gin-gonic/gin"
	"supermarket/app/models"
	"supermarket/app/response"
)

/**
 * 获取商品列表
 */
func GetGoods(c *gin.Context) {

	resGin := response.Gin{C: c}

	resGin.Success("invalid_option", nil)
}

/**
 * 获取商品详情
 */
func Detail(c *gin.Context) {
	// 响应
	resGin := response.Gin{C: c}
	// 响应成功
	resGin.Success("ok", nil)
}

// 猜你喜欢 tab
func GussULikeTab(c *gin.Context) {
	// 申明gridMenu
	var gussULikeModel models.GussULikeTab
	// store code
	storeCode := c.DefaultQuery("id","0000001")
	// fmt.Println(storeCode)
	// 获取数据
	gussULikeTabData := gussULikeModel.GetGussULikeTab(storeCode)
	// 响应
	resGin := response.Gin{C: c}
	// 响应成功
	resGin.Success("ok", gussULikeTabData)
}

// 猜你喜欢 商品
func GussULikeGoods(c *gin.Context) {
	// 申明gridMenu
	var gussULikeGoodsModel models.GussULikeGoods
	// 申明查询条件的map
	where := make(map[string]interface{})
	// 获取 门店编号为 0000001 的数据
	where["storeCode"] = "0000001"
	// 获取数据
	groupGoods := gussULikeGoodsModel.GetGussULikeGoods(where)
	// 响应
	resGin := response.Gin{C: c}
	// 响应成功
	resGin.Success("ok", groupGoods)
}
