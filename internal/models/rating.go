package model

import "time"

type Rating struct {
	Id        int        `json:"id" gorm:"column:id;primaryKey"`
	ProductId int        `json:"product_id" gorm:"column:product_id"`
	Product   Product    `gorm:"foreignKey:ProductId;constraint:OnUpdate:CasCade;OnDelete:SET NULL;"`
	Rating    float64    `json:"rating" gorm:"column:rating"`
	Comment   string     `json:"comment" gorm:"column:comment"`
	CreatedAt time.Time  `json:"created_at" gorm:"column:create_at"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"column:updated_at"`
	DeletedAt *time.Time `json:"-" gorm:"column:deleted_at"`
}
