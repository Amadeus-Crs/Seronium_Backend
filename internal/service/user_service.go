package service

import (
	"Seronium/internal/model"
	"Seronium/internal/repository"
	"crypto/rand"
	"encoding/base64"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo *repository.UserRepo
}

func (s *UserService) FindByID(userID uint64) (*model.User, error) {
	user := &model.User{}
	user, err := s.repo.FindByID(userID)
	return user, err
}

func NewUserService() *UserService {
	return &UserService{
		repo: &repository.UserRepo{},
	}
}

func (s *UserService) Register(username, password string) error {
	saltBytes := make([]byte, 16)
	_, err := rand.Read(saltBytes)
	if err != nil {
		return err
	}
	salt := base64.StdEncoding.EncodeToString(saltBytes)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password+salt), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user := &model.User{
		Username:     username,
		PasswordHash: string(hashedPassword),
		Salt:         salt,
	}
	return s.repo.Create(user)
}

func (s *UserService) Login(username, password string) (*model.User, error) {
	user, err := s.repo.FindByUsername(username)
	if err != nil {
		if err := s.Register(username, password); err != nil {
			return nil, err
		}
		return s.repo.FindByUsername(username)
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password+user.Salt)); err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserService) UpdateProfile(userID uint64, bio, avatarURL string) error {
	user := &model.User{
		ID:        userID,
		Bio:       bio,
		AvatarURL: avatarURL,
	}
	return s.repo.Update(user)
}
