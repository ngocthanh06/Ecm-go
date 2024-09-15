package model

type ProductTagRelationship struct {
	ProductId int `json:"product_id" gorm:"column:product_id;PrimaryKey"`
	TagId     int `json:"tag_id" gorm:"column:tag_id;PrimaryKey"`
}
