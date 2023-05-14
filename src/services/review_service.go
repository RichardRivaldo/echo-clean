package services

import (
	"echo-clean/src/configs"
	"echo-clean/src/models"

	"gorm.io/gorm"
)

func GetProductReview(id string) ([]models.Review, *gorm.DB) {
	data := []models.Review{}
	result := configs.DB.Table("reviews").Where("product_id = ?", id).
		Preload("Product").Preload("Member").Preload("LikeReviews").
		Find(&data)

	// Filling the count
	for i := range data {
		data[i].LikeCount = len(data[i].LikeReviews)
	}
	return data, result
}
