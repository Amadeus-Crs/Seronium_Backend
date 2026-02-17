package repository

import "Seronium/internal/model"

type UserRepo struct{}

func NewUserRepo() *UserRepo {
	return &UserRepo{}
}

func (r *UserRepo) Create(user *model.User) error {
	return DB.Create(user).Error
}

func (r *UserRepo) GetByUsername(username string) (*model.User, error) {
	var user model.User
	err := DB.Where("username=?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
