package repository

import "Seronium/internal/model"

type LikeRepo struct{}

func (r *LikeRepo) Create(like *model.Like) error {
	return DB.Create(like).Error
}

func (r *LikeRepo) Exists(userID uint64, targetType string, targetID uint64) bool {
	var count int64
	DB.Model(&model.Like{}).Where("user_id = ? AND target_type = ? AND target_id = ?", userID, targetType, targetID).Count(&count)
	return count > 0
}
