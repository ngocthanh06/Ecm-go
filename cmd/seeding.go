package cmd

import (
	"fmt"
	"github.com/bxcodec/faker/v3"
	"github.com/ngocthanh06/ecommerce/internal/database"
	model "github.com/ngocthanh06/ecommerce/internal/models"
	"github.com/ngocthanh06/ecommerce/pkg/utils"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/bcrypt"
	"log"
	"math/rand"
)

var (
	seedingCmd = &cobra.Command{
		Use:   "seeding",
		Short: "root is cli",
		Long:  "root is cli",
		Run: func(cmd *cobra.Command, args []string) {
			seedingDatabase()
		},
	}
)

// init
//
// Parameters:
//
// Returns:
func init() {
	rootCmd.AddCommand(seedingCmd)
}

// Seeding database
//
// Parameters:
//
// Returns:
func seedingDatabase() {
	// insert data in database
	users := usersSeeding()
	categories := categoriesSeeding()
	productTagsSeeding()
	for _, user := range users {
		for _, category := range categories {
			productSeeding(user.Id, category.Id)
		}
	}
}

// address Seeding
//
// Parameters:
// - userNumber: int
//
// Returns:
func addressSeeding(userNumber int) {
	address := []*model.Address{}
	// create address user
	for value := range userNumber {
		add := &model.Address{
			UserId:  value + 1,
			Address: faker.MacAddress(),
			City:    faker.MacAddress(),
		}

		address = append(address, add)
	}

	database.GetDb().Db.Create(&address)
}

// user Seeding
//
// Parameters:
//
// Returns:
// - []*model.Category
func usersSeeding() []*model.User {
	hash, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)

	users := []*model.User{
		{
			FirstName: "user",
			LastName:  "last",
			Email:     "user@gmail.com",
			Password:  string(hash),
			Phone:     "123456789",
			Role:      utils.Roles["user"],
		},
		{
			FirstName: "user",
			LastName:  "last",
			Email:     "admin@gmail.com",
			Password:  string(hash),
			Phone:     "123456788",
			Role:      utils.Roles["admin"],
		},
		{
			FirstName: "master",
			LastName:  "last",
			Email:     "master@gmail.com",
			Password:  string(hash),
			Phone:     "123456787",
			Role:      utils.Roles["master"],
		},
	}

	data := database.GetDb().Db.Create(&users)

	if data.Error != nil {
		log.Printf("insert users fails %s", data.Error)
	}

	// seeder address user
	addressSeeding(len(users))

	return users
}

// Categories Seeding
//
// Parameters:
//
// Returns:
// - []*model.Category
func categoriesSeeding() []*model.Category {
	categories := []*model.Category{
		{
			Name: "Cake & Milk",
		},
		{
			Name: "Coffes & Teas",
		},
		{
			Name: "Pet Foods",
		},
		{
			Name: "Vegetables",
		},
	}

	data := database.GetDb().Db.Create(&categories)

	if data.Error != nil {
		log.Printf("insert category fails %s", data.Error)
	}

	return categories
}

// Product tags Seeding
//
// Parameters:
//
// Returns:
func productTagsSeeding() {
	productTags := []*model.ProductTag{
		{
			Name:        "Recently added",
			Description: "Recently added",
		},
		{
			Name:        "Trending Products",
			Description: "Trending Products",
		},
		{
			Name:        "Top Selling",
			Description: "Top Selling",
		},
		{
			Name:        "Daily Best Sells",
			Description: "Daily Best Sells",
		},
	}

	data := database.GetDb().Db.Create(&productTags)

	if data.Error != nil {
		fmt.Printf("insert product tags fails: %v \n", data.Error)
	}
}

// Product Seeding
//
// Parameters:
// - userId: int
// - categoryId: int
//
// Returns:
func productSeeding(userId int, categoryId int) {
	products := []*model.Product{}

	for i := range 10 {
		product := &model.Product{
			Name:        faker.Name(),
			Description: faker.Word(),
			Price:       rand.Float64(),
			CategoryId:  categoryId,
			UserId:      userId,
		}

		// create product relationship
		_ = model.ProductTagRelationship{
			ProductId: product.Id,
			TagId:     rand.Intn(4),
		}

		products = append(products, product)

		i++
	}

	data := database.GetDb().Db.Create(&products)

	if data.Error != nil {
		fmt.Printf("insert product fails: %v \n", data.Error)
	}
}
