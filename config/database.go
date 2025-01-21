package config

import (
	"errors"
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

func GetDB() *gorm.DB {
	return db
}

func init() {
	_ = godotenv.Load(".env")

	databaseInisialisation := "" + os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD") + "@tcp(" + os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + ")/" + os.Getenv("DB_NAME") + "?charset=utf8mb4&parseTime=True&loc=Local"
	database, err := gorm.Open(mysql.Open(databaseInisialisation), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db = database

	err = db.AutoMigrate(&model.User{}, &model.Product{})

	if err != nil {
		panic(err.Error())
	}
	err = database.First(&model.User{}).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			Seeds(db)
		}
	}

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

	var products = []model.Product{
		model.Product{
			ID:   uuid.NewString(),
			Name: "Hijab Muslimah",
			Description: `Lorem ipsum dolor sit amet, consectetur adipiscing elit. Aliquam pretium arcu a ex rutrum malesuada. In sed urna sollicitudin, 
			interdum felis at, hendrerit erat. Fusce sodales vehicula turpis
			, ac placerat metus convallis ut. Aenean ut massa non turpis tristique iaculis in ut enim. Nunc porttitor metus eu volutpat mattis. Mauris.`,
			Stock:     200,
			SKU:       "PROD-001",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		model.Product{
			ID:   uuid.NewString(),
			Name: "Gamis",
			Description: `Lorem ipsum dolor sit amet, consectetur adipiscing elit. Aliquam pretium arcu a ex rutrum malesuada. In sed urna sollicitudin, 
			interdum felis at, hendrerit erat. Fusce sodales vehicula turpis
			, ac placerat metus convallis ut. Aenean ut massa non turpis tristique iaculis in ut enim. Nunc porttitor metus eu volutpat mattis. Mauris.`,
			Stock:     201,
			SKU:       "PROD-002",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
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

	return true
}
