package config

import (
	"os"
)

type (
	Config interface {
		Get(key string) string
	}

	configImpl struct {
	}
)

func (config *configImpl) Get(key string) string {
	return os.Getenv(key)
}

//func New(filenames ...string) Config {
//	err := godotenv.Load(filenames...)
//helper.
//}
