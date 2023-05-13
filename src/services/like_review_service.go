package services

import (
	"echo-clean/src/configs"
	"echo-clean/src/models"

	"gorm.io/gorm"
)

func Like(likeReview models.LikeReview) (models.LikeReview, *gorm.DB) {
	result := configs.DB.Create(&likeReview)
	return likeReview, result
}

func Unlike(likeReview models.LikeReview) *gorm.DB {
	result := configs.DB.Delete(&likeReview)
	return result
}
