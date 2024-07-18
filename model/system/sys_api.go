package system

import "time"

type SysApi struct {
	//global.GVA_MODEL
	ID        uint      `gorm:"primarykey" json:"ID"` // 主键ID
	CreatedAt time.Time // 创建时间
	UpdatedAt time.Time // 更新时间
	//DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`                        // 删除时间
	Path        string `json:"path" gorm:"comment:api路径"`             // api路径
	Description string `json:"description" gorm:"comment:api中文描述"`    // api中文描述
	ApiGroup    string `json:"apiGroup" gorm:"comment:api组"`          // api组
	Method      string `json:"method" gorm:"default:POST;comment:方法"` // 方法:创建POST(默认)|查看GET|更新PUT|删除DELETE
}
