package repository

import "Seronium/internal/model"

type CommentRepo struct {
}

func (r *CommentRepo) Create(comment *model.Comment) error {
	return DB.Create(comment).Error
}
