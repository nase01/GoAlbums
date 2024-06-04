package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CtmBasemModel struct {
	Id        string         `json:"id" gorm:"primaryKey;type:varchar(50);index"`
	CreatedAt time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"autoUpdateTime;index"`
	UpdatedBy string         `json:"updated_by" gorm:"column:updated_by; default:'system'"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

func (m *CtmBasemModel) BeforeCreate(tx *gorm.DB) (err error) {
	if m.Id == "" {
		m.Id = uuid.New().String()
	}
	return
}
