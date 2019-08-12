package routes

import (
	"github.com/gin-gonic/gin"
	"supermarket/app/http/controllers/banner"
	"supermarket/app/http/controllers/goods"
	"supermarket/app/http/controllers/gridMenu"
	"supermarket/app/http/controllers/groupGoods"
	"supermarket/app/http/middleware"
)

func InitRouter() *gin.Engine {
	// 初始化默认路由
	router := gin.Default()
	// 路由分组
	api := router.Group("api")
	{
		// 不使用中间价
		api.GET("banner/:type", banner.GetBanners)
		api.GET("grid_menu/:type", gridMenu.GetGridMenus)
		// 商品
		api.GET("goods", goods.GetGoods)
		// 团购商品
		api.GET("group_goods",groupGoods.GetGroupGoods)
		// 猜你喜欢
		api.GET("guss_u_like_tab",goods.GussULikeTab)
		api.GET("guss_u_like_goods",goods.GussULikeGoods)
		// 使用中间价
		api.Use(middleware.JWTAuth())
		api.GET("goods/:id", goods.Detail)
	}
	return router
}
