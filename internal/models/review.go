package model

import "time"

type Review struct {
	Id          int        `json:"id" gorm:"column:id;primaryKey"`
	Name        string     `json:"name" gorm:"column:name"`
	Description string     `json:"description" gorm:"column:description"`
	Price       float64    `json:"price" gorm:"column:price"`
	Stock       int        `json:"stock" gorm:"column:stock"`
	CategoryId  int        `json:"category_id" gorm:"column:category_id"`
	Category    Category   `gorm:"foreignKey:CategoryId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CreatedAt   time.Time  `json:"created_at" gorm:"column:create_at"`
	UpdatedAt   time.Time  `json:"updated_at" gorm:"column:updated_at"`
	DeletedAt   *time.Time `json:"-" gorm:"column:deleted_at"`
}
