package config

import (
	"context"
	"errors"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/cast"
	"gorm.io/gorm/logger"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/khaizbt/superindo-test/model"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB
var rdb *redis.Client

func GetDB() *gorm.DB {
	return db
}

func GetRedis() *redis.Client {
	return rdb
}

func init() {
	_ = godotenv.Load(".env")

	databaseInisialisation := "" + os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD") + "@tcp(" + os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + ")/" + os.Getenv("DB_NAME") + "?charset=utf8mb4&parseTime=True&loc=Local"
	database, err := gorm.Open(mysql.Open(databaseInisialisation), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		panic(err)
	}

	db = database

	err = db.AutoMigrate(&model.User{}, &model.Product{}, &model.ProductCategory{}, &model.Category{})

	if err != nil {
		panic(err.Error())
	}
	err = database.First(&model.User{}).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			Seeds(db)
		}
	}

	rClient := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDRESS"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       cast.ToInt(os.Getenv("REDIS_DB")),
	})

	_, err = rClient.Ping(context.Background()).Result()

	if err != nil {
		panic(err.Error())
	}

	rdb = rClient
}

func Seeds(db *gorm.DB) bool { //https://gorm.io/docs/migration.html
	passwordHash, err := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.MinCost)

	if err != nil {
		panic(err.Error())
	}

	var users = []model.User{
		model.User{
			ID:       uuid.NewString(),
			Username: "user_superindo",
			Email:    "user@superindo.com",
			Password: string(passwordHash),
		},
		model.User{
			ID:       uuid.NewString(),
			Username: "user_superindo1",
			Email:    "user2@superindo.com",
			Password: string(passwordHash),
		},
	}

	var description = `Lorem ipsum dolor sit amet, consectetur adipiscing elit. Aliquam pretium arcu a ex rutrum malesuada. In sed urna sollicitudin,
		interdum felis at, hendrerit erat. Fusce sodales vehicula turpis
	, ac placerat metus convallis ut. Aenean ut massa non turpis tristique iaculis in ut enim. Nunc porttitor metus eu volutpat mattis. Mauris`
	var products = []model.Product{
		model.Product{
			ID:          "fbde8f5f-cce9-471e-addc-f0f019a7f755",
			Name:        "Beef Wagyu A5 500gr",
			Description: &description,
			Stock:       5,
			SKU:         "B001",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		model.Product{
			ID:          "9fd18521-6358-44de-a426-901572dc9403",
			Name:        "Kangkung Segar Dieng",
			Description: &description,
			Stock:       201,
			SKU:         "S001",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
	}

	var categories = []model.Category{
		model.Category{
			ID:          "f522ae24-cc07-483b-97bc-7895468f6146",
			Name:        "Sayuran",
			Description: "Lorem ipsum dolor sit amet",
		},
		model.Category{
			ID:          "b0ea46a9-f142-4d99-a5eb-926061bfe145",
			Name:        "Buah",
			Description: "Lorem ipsum dolor sit amet",
		},
		model.Category{
			ID:          "5edbf5f9-15e6-4202-866d-656b515d5419",
			Name:        "Snack",
			Description: "Lorem ipsum dolor sit amet",
		},
		model.Category{
			ID:          "5edbf5f9-15e6-4202-866d-656b515d5420",
			Name:        "Protein",
			Description: "Lorem ipsum dolor sit amet",
		},
	}

	var productCategories = []model.ProductCategory{
		model.ProductCategory{
			ID:         uuid.NewString(),
			CategoryID: "5edbf5f9-15e6-4202-866d-656b515d5420",
			ProductID:  "fbde8f5f-cce9-471e-addc-f0f019a7f755",
		},
		model.ProductCategory{
			ID:         uuid.NewString(),
			CategoryID: "f522ae24-cc07-483b-97bc-7895468f6146",
			ProductID:  "9fd18521-6358-44de-a426-901572dc9403",
		},
		model.ProductCategory{
			ID:         uuid.NewString(),
			CategoryID: "5edbf5f9-15e6-4202-866d-656b515d5420",
			ProductID:  "9fd18521-6358-44de-a426-901572dc9403",
		},
	}

	//TODO seed products

	err = db.Migrator().DropTable(&model.User{})
	if err != nil {
		panic(err.Error())
	}

	err = db.AutoMigrate(&model.User{}, &model.Product{})

	if err != nil {
		panic(err.Error())
	}

	for i, _ := range users {
		err = db.Debug().Model(&model.User{}).Create(&users[i]).Error

		if err != nil {
			panic("Migration Failed")
		}
	}

	for i, _ := range products {
		err = db.Debug().Model(&model.Product{}).Create(&products[i]).Error

		if err != nil {
			panic("Migration Failed")
		}
	}

	for i, _ := range categories {
		err = db.Debug().Model(&model.Category{}).Create(&categories[i]).Error
		if err != nil {
			panic("Seeding Failed")
		}
	}

	for i, _ := range productCategories {
		err = db.Debug().Model(&model.ProductCategory{}).Create(&productCategories[i]).Error
		if err != nil {
			panic("Seeding Failed")
		}
	}
	return true
}
