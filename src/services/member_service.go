package services

import (
	"echo-clean/src/configs"
	"echo-clean/src/models"

	"gorm.io/gorm"
)

func GetAllMembers() ([]models.Member, error) {
	var members []models.Member
	result := configs.DB.Find(&members)

	return members, result.Error
}

func CreateNewMember(member models.Member) (models.Member, *gorm.DB) {
	result := configs.DB.Create(&member)
	return member, result
}

func UpdateMember(id string, member models.Member) *gorm.DB {
	result := configs.DB.Where("member_id = ?", id).Updates(member)
	return result
}

func DeleteMember(id string) *gorm.DB {
	result := configs.DB.Where("member_id = ?", id).Delete(&models.Member{})
	return result
}
