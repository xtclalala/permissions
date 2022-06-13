package system

type SysFile struct {
	BaseID
	Name string `json:"name" gorm:"comment:文件名称;"`
	Type string `json:"type" gorm:"comment:文件类型;"`
	Path int64  `json:"path" gorm:"comment:文件存储路径;"`
}
