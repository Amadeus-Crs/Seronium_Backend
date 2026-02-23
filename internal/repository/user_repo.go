package repository

import "Seronium/internal/model"

type UserRepo struct{}

func (r *UserRepo) Create(user *model.User) error {
	return DB.Create(user).Error
}

func (r *UserRepo) FindByUsername(username string) (*model.User, error) {
	var user model.User
	err := DB.Where("username = ?", username).First(&user).Error
	return &user, err
}
func (r *UserRepo) FindByID(id uint64) (*model.User, error) {
	var user model.User
	err := DB.Where("id = ?", id).First(&user).Error
	return &user, err
}

func (r *UserRepo) Update(user *model.User) error {
	return DB.Save(user).Error
}
