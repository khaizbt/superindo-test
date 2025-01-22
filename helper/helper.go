package helper

import (
	"crypto/rand"
	"errors"
	"fmt"
	"github.com/khaizbt/superindo-test/config"
	"github.com/khaizbt/superindo-test/model"
)

func GenerateNumber(length int) (string, error) {
	var chars = "0123456789"
	buffer := make([]byte, length)
	_, err := rand.Read(buffer)
	if err != nil {
		return "0", err
	}

	charsLength := len(chars)
	for i := 0; i < length; i++ {
		buffer[i] = chars[int(buffer[i])%charsLength]
	}

	return string(buffer), nil
}

func GetCategory(query []string) error {
	db := config.GetDB()

	for _, category := range query {
		err := db.Where("id = ?", category).First(&model.Category{}).Error

		if err != nil {
			return errors.New(fmt.Sprint("category does not exist: ", category))
		}
	}

	return nil

}
