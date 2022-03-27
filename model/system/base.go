package system

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

// BaseUUID uuid模板
type BaseUUID struct {
	ID         uuid.UUID `json:"id" gorm:"type:varchar(36);primary_key"`
	CreateTime time.Time
	UpdateTime time.Time
}

// BeforeCreate 创建时添加uuid
func (bUuid *BaseUUID) BeforeCreate(tx *gorm.DB) (err error) {
	bUuid.ID = uuid.New()
	bUuid.CreateTime = time.Now()
	bUuid.UpdateTime = time.Now()
	return
}

// AfterUpdate 更新时添加更新时间
func (bUuid *BaseUUID) AfterUpdate(tx *gorm.DB) (err error) {
	bUuid.UpdateTime = time.Now()
	return
}

// BaseID 自增id模板
type BaseID struct {
	ID int `gorm:"primary_key"`
}
