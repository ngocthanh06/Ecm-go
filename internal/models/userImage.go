package model

import "time"

type UserImage struct {
	Id        int        `json:"id" gorm:"column:id; primaryKey"`
	UserId    int        `json:"user_id" gorm:"column:user_id;not null;"`
	ImageId   int        `json:"image_id" gorm:"column:image_id;not null;"`
	CreatedAt time.Time  `json:"created_at" gorm:"column:create_at"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"column:updated_at"`
	DeletedAt *time.Time `json:"-" gorm:"column:deleted_at"`
}
