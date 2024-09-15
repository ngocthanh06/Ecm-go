package model

import "time"

type Image struct {
	Id        int        `json:"id" gorm:"column:id;primaryKey"`
	Url       string     `json:"url" gorm:"column:url"`
	Type      string     `json:"type" gorm:"column:type;not null"`
	CreatedAt time.Time  `json:"created_at" gorm:"column:create_at"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"column:updated_at"`
	DeletedAt *time.Time `json:"-" gorm:"column:deleted_at"`
}
