package model

import "time"

type ProductTag struct {
	Id          int        `json:"id" gorm:"column:id;primaryKey"`
	Name        string     `json:"name" gorm:"column:name"`
	Description string     `json:"description" gorm:"column:description"`
	CreatedAt   time.Time  `json:"created_at" gorm:"column:create_at"`
	UpdatedAt   time.Time  `json:"updated_at" gorm:"column:updated_at"`
	DeletedAt   *time.Time `json:"-" gorm:"column:deleted_at"`
}
