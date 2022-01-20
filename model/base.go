package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

// CRUD 时间模板
type CRUD struct {
	CreateTime time.Time
	UpdateTime time.Time
}

// BeforeCreate 创建时添加创建时间
func (crud *CRUD) BeforeCreate(tx *gorm.DB) (err error) {
	crud.CreateTime = time.Now()
	return
}

// AfterUpdate 更新时添加更新时间
func (crud *CRUD) AfterUpdate(tx *gorm.DB) (err error) {
	crud.UpdateTime = time.Now()
	return
}

// BaseUUID uuid模板
type BaseUUID struct {
	ID uuid.UUID `json:"id" gorm:"type:varchar(36);primary_key"`
	CRUD
}

// BeforeCreate 创建时添加uuid
func (bUuid *BaseUUID) BeforeCreate(tx *gorm.DB) (err error) {
	bUuid.ID = uuid.New()
	return
}

// BaseID 自增id模板
type BaseID struct {
	ID uint `gorm:"primary_key"`
	CRUD
}
