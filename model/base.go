package model

import (
	"github.com/google/uuid"
	"time"
)

type baseTime struct {
	Id         uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	CreateDate time.Time `gorm:`
	Deleted    bool      ``
}

type baseNum struct {
}
