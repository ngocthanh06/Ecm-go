package database

import (
	"github.com/ngocthanh06/ecommerce/internal/models"
	"gorm.io/gorm"
	"log"
)

func AutoMigration(db *gorm.DB) {
	var modelTables = []interface{}{
		&model.User{},
		&model.Category{},
		&model.Product{},
		&model.ProductTag{},
		&model.ProductTagRelationship{},
		&model.Coupon{},
		&model.Address{},
		&model.Order{},
		&model.Payment{},
		&model.OrderItem{},
		&model.Review{},
		&model.ShoppingCart{},
		&model.Image{},
		&model.CategoryImage{},
		&model.ProductImage{},
		&model.UserImage{},
	}

	if err := db.AutoMigrate(modelTables...); err != nil {
		log.Fatalf("Migrate error: %v", err)
	}
}
