package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BaseSoftDel struct {
	ID        uuid.UUID  `gorm:"type:uuid;primary_key"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`
}

func (b *BaseSoftDel) BeforeCreate(tx *gorm.DB) (err error) {
	b.ID = uuid.New()
	return
}
