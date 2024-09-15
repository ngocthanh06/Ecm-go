package model

type Payment struct {
	Id            int    `json:"id" gorm:"column:id;primaryKey"`
	OrderId       int    `json:"order_id" gorm:"column:order_id"`
	Order         *Order `gorm:"foreignKey:OrderId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	PaymentMethod int    `json:"payment_method" gorm:"column:payment_method"`
	PaymentStatus int    `json:"payment_status" gorm:"column:payment_status"`
	PaymentDate   int    `json:"payment_date" gorm:"column:payment_date"`
}
