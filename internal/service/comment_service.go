package service

import (
	"Seronium/internal/model"
	"Seronium/internal/repository"
	"strconv"
)

type CommentService struct {
	repo *repository.CommentRepo
}

func NewCommentService() *CommentService {
	return &CommentService{repo: &repository.CommentRepo{}}
}

func (s *CommentService) Create(userID uint64, targetType string, targetID uint64, content string) error {
	comment := &model.Comment{UserID: userID, TargetType: targetType, TargetID: targetID, Content: content}
	if err := s.repo.Create(comment); err != nil {
		return err
	}
	if targetType == "post" {
		repository.RDB.ZIncrBy(repository.CTX, "post:hot", 3, strconv.FormatUint(targetID, 10))
	}
	return nil
}
