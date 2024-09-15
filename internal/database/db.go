package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"sync"
)

type envDbConfig struct {
	host     string
	port     string
	username string
	password string
	database string
}

type Database struct {
	Db *gorm.DB
}

var databaseInstance *Database
var once sync.Once

func setDatabase(db *gorm.DB) {
	databaseInstance = &Database{
		Db: db,
	}
}

func GetDb() *Database {
	return databaseInstance
}

func newEnvDbConfig() *envDbConfig {
	return &envDbConfig{
		host:     os.Getenv("DB_HOST"),
		port:     os.Getenv("PORT"),
		username: os.Getenv("DB_USER"),
		password: os.Getenv("DB_PASS"),
		database: os.Getenv("DB_NAME"),
	}
}

func getDnsDatabaseSqlite(config *envDbConfig) string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		config.host,
		config.username,
		config.password,
		config.database,
		config.port,
	)
}

func ConnectionDatabase() {
	once.Do(func() {
		dbConfig := newEnvDbConfig()
		dns := getDnsDatabaseSqlite(dbConfig)

		dbConnect, err := gorm.Open(postgres.Open(dns), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})

		if err != nil {
			log.Fatalf("Could not connect to the database: %v", err)
		}

		fmt.Println("Connection database successfully!")

		setDatabase(dbConnect)

		AutoMigration(GetDb().Db)
	})
}
