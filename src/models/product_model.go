package models

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Product struct {
	ProductID   string          `json:"product_id,omitempty" gorm:"primaryKey"`
	ProductName string          `json:"product_name" validate:"required"`
	Price       decimal.Decimal `json:"price" validate:"required"`
}

func (p *Product) BeforeCreate(tx *gorm.DB) (err error) {
	p.ProductID = uuid.New().String()
	return
}
