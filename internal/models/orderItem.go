package model

type OrderItem struct {
	Id        int     `json:"id" gorm:"column:id;primaryKey"`
	OrderId   int     `json:"order_id" gorm:"column:order_id"`
	Order     Order   `gorm:"foreignKey:OrderId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	ProductId int     `json:"product_id" gorm:"column:product_id"`
	Quantity  int     `json:"quantity" gorm:"column:quantity"`
	Price     float64 `json:"price" gorm:"column:price"`
}
