package repository

import "Seronium/internal/model"

type CollectionRepo struct{}

func (r *CollectionRepo) Create(collection *model.Collection) error {
	return DB.Create(collection).Error
}

func (r *CollectionRepo) Exists(userID, postID uint64) bool {
	var count int64
	DB.Model(&model.Collection{}).Where("user_id = ? AND post_id = ?", userID, postID).Count(&count)
	return count > 0
}
