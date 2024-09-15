package model

import "time"

type Coupon struct {
	Id        int        `json:"id" gorm:"column:id;primaryKey"`
	Code      string     `json:"code" gorm:"code"`
	Discount  float64    `json:"discount" gorm:"discount"`
	ValidForm *time.Time `json:"valid_form" gorm:"valid_form"`
	ValidTo   *time.Time `json:"valid_to" gorm:"valid_to"`
}
