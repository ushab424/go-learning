package service

import (
	"day19tier2/repository"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetUser(id int) (*repository.User, error) {
	return s.repo.GetByID(id)
}

func (s *UserService) Register(name, email, password string, age int) (int, error) {
	// validation
	if name == "" || email == "" || password == "" {
		return 0, errors.New("missing required fields")
	}

	// hash password
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return 0, err
	}

	// save in to DB
	UserID, err := s.repo.Create(name, email, string(hash), age)
	return UserID, err
}

func (s *UserService) Login(email, password string) (*repository.User, error) {
	// search user
	user, hash, err := s.repo.GetByEmail(email)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	// check password
	if err := bcrypt.CompareHashAndPassword([]byte(*hash), []byte(password)); err != nil {
		return nil, errors.New("invalid credentials")
	}

	return user, nil
}

func (s *UserService) DeleteUser(id int) error {
	return s.repo.Delete(id)
}
