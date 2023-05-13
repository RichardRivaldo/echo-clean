package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Review struct {
	ReviewID   string `json:"review_id" gorm:"primaryKey"`
	MemberID   string `json:"member_id"`
	ProductID  string `json:"product_id"`
	DescReview string `json:"desc_review"`
}

func (r *Review) BeforeCreate(tx *gorm.DB) (err error) {
	r.ReviewID = uuid.New().String()
	return
}
