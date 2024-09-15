package model

type Address struct {
	Id         int    `json:"id" gorm:"column:id;primaryKey"`
	UserId     int    `json:"user_id" gorm:"column:user_id"`
	User       User   `gorm:"foreignKey:UserId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	Address    string `json:"address" gorm:"column:address"`
	City       string `json:"city" gorm:"column:city"`
	PostalCode string `json:"-" gorm:"column:postal_code"`
	Country    string `json:"-" gorm:"country"`
}
