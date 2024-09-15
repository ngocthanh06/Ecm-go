package model

import "time"

type Order struct {
	Id                int        `json:"id" gorm:"column:id;primaryKey"`
	UserId            int        `json:"user_id" gorm:"column:user_id"`
	User              User       `gorm:"foreignKey:UserId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	TotalPrice        float64    `json:"total_price" gorm:"column:total_price"`
	Status            string     `json:"status" gorm:"column:status"`
	ShippingAddressId int        `json:"shipping_address_id" gorm:"column:shipping_address_id"`
	CreatedAt         time.Time  `json:"created_at" gorm:"column:create_at"`
	UpdatedAt         time.Time  `json:"updated_at" gorm:"column:updated_at"`
	DeletedAt         *time.Time `json:"-" gorm:"column:deleted_at"`
}
