package groupGoods

import (
	"github.com/gin-gonic/gin"
	"supermarket/app/models"
	"supermarket/app/response"
)

/**
 * 获取商品列表
 */
func GetGroupGoods(c *gin.Context) {
	// 申明gridMenu
	var groupGoodsModel models.GroupGoods
	// 申明查询条件的map
	where := make(map[string]interface{})
	// 获取 门店编号为 0000001 的数据
	where["storeCode"] = "0000001"
	// 获取数据
	groupGoods := groupGoodsModel.GetGroupGoodsByCondition(where)
	// 响应
	resGin := response.Gin{C: c}
	// 响应成功
	resGin.Success("ok", groupGoods)
}

/**
 * 获取商品详情
 */
func GroupDetail(c *gin.Context) {
	resGin := response.Gin{C: c}

	resGin.Success("page_not_found", nil)
}
