package model

type ShoppingCart struct {
	Id        int     `json:"id" gorm:"column:id;primaryKey"`
	UserId    int     `json:"user_id" gorm:"column:user_id"`
	ProductId int     `json:"product_id" gorm:"column:product_id"`
	User      User    `gorm:"foreignKey:UserId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Product   Product `gorm:"foreignKey:ProductId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Quantity  int     `json:"quantity" gorm:"column:quantity"`
}
