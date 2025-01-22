package service

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/khaizbt/superindo-test/config"
	"log"
	"time"

	"github.com/khaizbt/superindo-test/entity"
	"github.com/khaizbt/superindo-test/model"
	"github.com/khaizbt/superindo-test/repository"
	"golang.org/x/crypto/bcrypt"
)

type (
	service struct {
		repository repository.UserRepository
	}

	UserService interface {
		CreateUser(input entity.DataUserInput) (bool, error)
		GetUserById(Id string) (model.User, error)
		Login(input entity.LoginInput) (model.User, error)
	}

	OutputHistoryBalance struct {
		Date        time.Time `json:"date"`
		Type        string    `json:"type"`
		DebitCredit int       `json:"debit_credit"`
		Balance     int       `json:"balance"`
	}
)

func NewUserService(repository repository.UserRepository) *service {
	return &service{repository}
}

func (s *service) CreateUser(input entity.DataUserInput) (bool, error) {
	cekUser, err := s.repository.FindByEmail(input.Email)

	if cekUser.ID != "" {
		return false, errors.New("Email has been registered")
	}

	cekUsername, err := s.repository.FindByUsername(input.Username)

	if cekUsername.ID != "" {
		return false, errors.New("Username has been taken")
	}

	user := model.User{
		Email:    input.Email,
		Username: input.Username,
	}

	password, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)

	if err != nil {
		return false, err
	}

	user.Password = string(password)

	_, err = s.repository.CreateUser(user)

	if err != nil {
		return false, err
	}

	return true, nil
}

func (s *service) GetUserById(id string) (model.User, error) {
	var rdb = config.GetRedis()
	user, err := s.repository.FindByID(id)

	if err != nil {
		return user, err
	}

	user.Password = ""
	userJson, err := json.Marshal(user)

	err = rdb.Set(context.Background(), user.ID, userJson, 2*time.Hour).Err()

	if err != nil {
		log.Println(err)
		return user, err
	}

	return user, nil
}

func (s *service) GetUserByUsername(username string) (model.User, error) {
	user, err := s.repository.FindByUsername(username)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *service) Login(input entity.LoginInput) (model.User, error) {
	email := input.Email
	password := input.Password

	user, err := s.repository.FindByEmail(email)

	if err != nil {
		return user, errors.New("user or password incorrect")
	}

	if user.ID == "" {
		return user, errors.New("email or password incorrect")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		return user, errors.New("email or password incorrect")
	}

	return user, nil
}
