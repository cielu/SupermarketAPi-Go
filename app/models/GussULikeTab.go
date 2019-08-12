package models

type GussULikeTab struct {
	TabName  string `json:"tabName" gorm:"size:255"`
	TabDesc  string `json:"tabDesc" gorm:"size:255"`
	Handwork string `json:"handwork" gorm:"size:20"`
}

// 获取表名
func (GussULikeTab) TableName() string {
	return "guss_u_like_tab"
}

// 获取groupGoods 菜单
func (groupGoods GussULikeTab) GetGussULikeTab(storeCode string) (gussULikeTabArr []GussULikeTab) {
	// 获取数据
	db.Select("*").Where("store_code = ?",storeCode).Limit(8).Order("sort desc").Find(&gussULikeTabArr)
	return
}
