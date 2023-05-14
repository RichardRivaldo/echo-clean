package models

import (
	"echo-clean/src/types"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Member struct {
	MemberID  string          `json:"member_id,omitempty" gorm:"primaryKey"`
	Username  string          `json:"username" validate:"required"`
	Gender    types.Gender    `json:"gender" validate:"required"`
	SkinType  types.SkinType  `json:"skin_type" validate:"required"`
	SkinColor types.SkinColor `json:"skin_color" validate:"required"`
}

func (m *Member) BeforeCreate(tx *gorm.DB) (err error) {
	m.MemberID = uuid.New().String()
	return
}
