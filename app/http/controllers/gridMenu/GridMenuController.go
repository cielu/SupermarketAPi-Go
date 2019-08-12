package gridMenu

import (
	"github.com/gin-gonic/gin"
	"supermarket/app/models"
	"supermarket/app/response"
)

/**
 * 获取商品列表
 */
func GetGridMenus(c *gin.Context) {
	// 申明gridMenu
	var GridMenusModel models.GridMenu
	// 获取请求参数
	gridMenuType := c.Param("type")
	// 判断获取的gridMenu类型
	gridMenus := GridMenusModel.GetGridMenuByType(gridMenuType)
	// 响应
	resGin := response.Gin{C: c}

	resGin.Success("OK", gridMenus)

}
