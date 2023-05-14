package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Review struct {
	ReviewID    string       `json:"review_id,omitempty" gorm:"primaryKey"`
	MemberID    string       `json:"member_id,omitempty"`
	ProductID   string       `json:"product_id,omitempty"`
	DescReview  string       `json:"desc_review"`
	Product     Product      `json:"product"`
	Member      Member       `json:"member"`
	LikeReviews []LikeReview `json:"like_reviews"`
	LikeCount   int          `json:"like_count"`
}

func (r *Review) BeforeCreate(tx *gorm.DB) (err error) {
	r.ReviewID = uuid.New().String()
	return
}
