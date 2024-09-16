package model

import "time"

type User struct {
	Id        int        `json:"id" gorm:"primaryKey;autoIncrement"`
	FirstName string     `json:"first_name" gorm:"column:first_name"`
	LastName  string     `json:"last_name" gorm:"column:last_name"`
	Email     string     `json:"email" gorm:"column:email;unique"`
	Password  string     `json:"Password" gorm:"column:password"`
	Phone     string     `json:"phone" gorm:"column:phone;unique"`
	Role      int        `json:"role" gorm:"column:role"`
	CreatedAt time.Time  `json:"created_at" gorm:"column:create_at"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"column:updated_at"`
	DeletedAt *time.Time `json:"-" gorm:"column:deleted_at"`
}

func (User) TableName() string {
	return "users"
}
