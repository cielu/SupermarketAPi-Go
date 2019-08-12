package models

type GussULikeGoods struct {
	GoodsCode          string  `json:"goodsCode" gorm:"size:20"`
	GoodsName          string  `json:"goodsName" gorm:"size:255"`
	GoodsPrice         float32 `json:"goodsPrice,string" gorm:"size:255"`
	GoodsSalePoint     string  `json:"goodsSalePoint" gorm:"size:255"`
	GoodsSalePrice     float32 `json:"goodsSalePrice,string" gorm:"size:255"`
	GoodsSalePriceType int8    `json:"goodsSalePriceType,string" gorm:"size:255"`
	GoodsSource        string  `json:"goodsSource"`
	GoodsStartNum      uint16  `json:"goodsStartNum,string"`
	GoodsStoreCode     string  `json:"goodsStoreCode"`
	GoodsCategory      string  `json:"goodsCategory"`
	GoodsCoverUrl      string  `json:"goodsCoverUrl" gorm:"size:255"`
	MerchantCode       string  `json:"merchantCode" gorm:"column:merchant_code"`
	SupplierCode       string  `json:"supplierCode"`
	Handwork           string  `json:"handwork"`
	IsPreSale          uint8   `json:"isPreSale,string"`
}

// 获取表名
func (GussULikeGoods) TableName() string {
	return "goods"
}

// 获取groupGoods 菜单
func (groupGoods GussULikeGoods) GetGussULikeGoods(where map[string]interface{}) (mapData map[string]interface{}) {
	// 申明 变量
	var (
		whereStr       string
		gussULikeGoods []GussULikeGoods
	)
	// 获取有效的团购商品
	sql := db.Select("*").Order("sort desc")
	// 店铺编码
	if storeCode, ok := where["storeCode"]; ok {
		whereStr += "goods_store_code = '" + storeCode.(string) + "' AND"
	}
	mapData = Pagination(sql,where,&gussULikeGoods)
	// map data
	return
}
