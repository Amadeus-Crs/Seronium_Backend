package repository

import (
	"Seronium/internal/model"
)

type PostRepo struct{}

func (r *PostRepo) Create(post *model.Post) error {
	return DB.Create(post).Error
}

func (r *PostRepo) FindByID(id uint64) (*model.Post, error) {
	var post model.Post
	err := DB.First(&post, id).Error
	return &post, err
}

func (r *PostRepo) Update(post *model.Post) error {
	return DB.Save(post).Error
}

func (r *PostRepo) Delete(id uint64) error {
	return DB.Where("id = ?", id).Delete(&model.Post{}).Error
}

func (r *PostRepo) List(offset, limit int) ([]model.Post, error) {
	var posts []model.Post
	err := DB.Offset(offset).Limit(limit).Find(&posts).Error
	return posts, err
}
