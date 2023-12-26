package service

import "gorm.io/gorm"

type ProductService struct {
	db *gorm.DB
}
