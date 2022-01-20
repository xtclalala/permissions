package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type CRUD struct {
	CreateTime time.Time
	UpdateTime time.Time
}

func (crud *CRUD) BeforeCreate(tx *gorm.DB) (err error) {
	crud.CreateTime = time.Now()
	return
}

func (crud *CRUD) AfterUpdate(tx *gorm.DB) (err error) {
	crud.UpdateTime = time.Now()
	return
}

type BaseUUID struct {
	ID uuid.UUID `json:"id" gorm:"type:varchar(36);primary_key"`
	CRUD
}

func (bUuid *BaseUUID) BeforeCreate(tx *gorm.DB) (err error) {
	bUuid.ID = uuid.New()
	return
}

type BaseID struct {
	ID uint `gorm:"primary_key"`
	CRUD
}
