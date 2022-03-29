package user

import "golang.org/x/crypto/bcrypt"

type Service interface {
	RegisterUser(input InputRegister) (User, error)
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
