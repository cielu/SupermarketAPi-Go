package models

type GridMenu struct {
	GridEleId uint   `json:"gridEleId" gorm:"primary_key;AUTO_INCREMENT"`
	EleName   string `json:"eleName" gorm:"size:255"`
	EleType   string `json:"eleType" gorm:"size:255"`
	ImgUrl    string `json:"imgUrl" gorm:"size:255"`
	Scene     string `json:"scene"`
	TargetUrl string `json:"targetUrl"`
	TargetId  uint   `json:"targetId"`
	ShopId    uint   `json:"shopId,-"`
	Sort      uint   `json:"sort"`
}

// 获取表名
func (GridMenu) TableName() string {
	return "grid_menus"
}

// 获取grid 菜单
func (grid GridMenu) GetGridMenuByType(_type string) (gridMenus []GridMenu) {
	var scene string
	// 初始化map
	switch _type {
	case "mobileIndex": // 小程序首页
		scene = "mobileIndex"
		break
	default:
		scene = "mobileIndex"
	}
	// 获取数据
	db.Table("grid_menus").
		Select("ele_name,ele_type,img_url,target_url,target_id").
		Where("scene = ?", scene).
		Limit(20).Order("sort desc").Find(&gridMenus)
	return
}
