package models

import (
	"time"
)

type GroupGoods struct {
	// Id           uint    `json:"-" gorm:"primary_key;AUTO_INCREMENT;-"`
	ActivityId   string  `json:"activityId" gorm:"size:20"`
	GoodsCode    string  `json:"goodsCode" gorm:"size:20"`
	GoodsName    string  `json:"goodsName" gorm:"size:255"`
	GoodsPic     string  `json:"goodsPic" gorm:"size:255"`
	GroupNum     uint16  `json:"groupNum,string"`
	MerchantCode string  `json:"merchantCode" gorm:"column:merchant_code"`
	GroupPrice   float32 `json:"groupPrice,string"`
}

// 获取表名
func (GroupGoods) TableName() string {
	return "group_goods"
}

// 获取groupGoods 菜单
func (groupGoods GroupGoods) GetGroupGoodsByCondition(where map[string]interface{}) (groupGoodsArr []GroupGoods) {
	// 申明变量
	var whereStr string
	// 获取有效的团购商品
	sql := db.Select("activity_id,goods_code,goods_name,goods_pic,group_num,merchant_code,group_price").Order("sort desc")
	// 店铺编码
	if storeCode, ok := where["storeCode"]; ok {
		whereStr += "store_code = '" + storeCode.(string) + "' AND"
	}
	// 获取当前时间
	nowTime := time.Now().Format("2006-01-02 15:04:05")
	// 获取有效数据
	sql = sql.Where(whereStr + " start_time < ? AND end_time > ?", nowTime, nowTime)
	// 获取数据
	sql.Limit(20).Order("sort desc").Find(&groupGoodsArr)

	return
}
