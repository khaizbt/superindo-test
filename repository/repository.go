package repository

import (
	"github.com/khaizbt/superindo-test/config"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository() *repository {
	return &repository{config.GetDB()}
}
