package user

import (
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterUser(input InputRegister) (User, error)
	LoginUser(input InputLogin) (User, error)
	IsDuplicateEmail(email string) (bool, error)
	SaveAvatar(Id int, fileLocation string) (User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) RegisterUser(input InputRegister) (User, error) {

	user := User{
		Email:      input.Email,
		Name:       input.Name,
		Occupation: input.Occupation,
		Role:       "user",
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}
	user.Password = string(hash)

	newUser, err := s.repository.Save(user)
	if err != nil {
		return newUser, nil
	}

	return newUser, nil
}

func (s *service) LoginUser(input InputLogin) (User, error) {
	email := input.Email
	password := input.Password

	user, err := s.repository.FindByEmail(email)

	if err != nil {
		return user, nil
	}

	if user.Id == 0 {
		return user, errors.New("user not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		return user, errors.New("invalid email/password")
	}

	return user, nil

}

func (s *service) IsDuplicateEmail(email string) (bool, error) {
	user, err := s.repository.FindByEmail(email)

	fmt.Println("cek email user", user)

	if err != nil {
		return false, err
	}

	if user.Id != 0 {
		return true, nil
	}

	return false, nil
}

func (s *service) SaveAvatar(Id int, fileLocation string) (User, error) {

	user, err := s.repository.FindById(Id)

	if err != nil {
		return user, err
	}

	user.Avatar = fileLocation

	userUpdate, err := s.repository.Update(user)

	if err != nil {
		return user, err
	}

	return userUpdate, nil

}
