package system

type SysFile struct {
	BaseUUID
	Name string `json:"name" gorm:"comment:文件名称;"`
	Type string `json:"type" gorm:"comment:文件类型;"`
	Path string `json:"path" gorm:"comment:文件存储路径;"`
}
