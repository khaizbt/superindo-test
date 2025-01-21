package repository

import (
	"github.com/khaizbt/superindo-test/model"
)

type (
	UserRepository interface {
		CreateUser(user model.User) (model.User, error)
		FindByID(ID string) (model.User, error)
		FindByEmail(Email string) (model.User, error)
		FindByUsername(Username string) (model.User, error)
	}
)

func (r *repository) CreateUser(user model.User) (model.User, error) {
	err := r.db.Create(&user).Error

	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *repository) FindByID(ID string) (model.User, error) {
	var user model.User

	err := r.db.Where("id = ?", ID).First(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) FindByEmail(Email string) (model.User, error) {
	var user model.User

	err := r.db.Where("email = ?", Email).First(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) FindByUsername(Username string) (model.User, error) {
	var user model.User

	err := r.db.Where("username = ?", Username).First(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}
