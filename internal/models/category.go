package model

import "time"

type Category struct {
	Id        int        `json:"id" gorm:"column:id;primaryKey"`
	Name      string     `json:"name" gorm:"column:name;unique"`
	ParentId  *int       `json:"parentId" gorm:"column:parent_id"`
	Parent    *Category  `gorm:"foreignKey:ParentId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CreatedAt time.Time  `json:"created_at" gorm:"column:create_at"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"column:updated_at"`
	DeletedAt *time.Time `json:"-" gorm:"column:deleted_at"`
}
