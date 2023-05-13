package models

type LikeReview struct {
	ReviewID string `json:"review_id" gorm:"primaryKey"`
	MemberID string `json:"member_id" gorm:"primaryKey"`
}
