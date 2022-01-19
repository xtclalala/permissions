package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type CRUD struct {
	CreateTime time.Time      `gorm:"autoCreateTime:milli"`
	UpdateTime time.Time      `gorm:"autoUpdateTime:milli"`
	Deleted    gorm.DeletedAt `gorm:"index" json:"deleted"`
}

type baseUUID struct {
	Id uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	CRUD
}

type baseID struct {
	Id uint `gorm:"primary_key"`
	CRUD
}
