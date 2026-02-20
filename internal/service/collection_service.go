package service

import (
	"Seronium/internal/model"
	"Seronium/internal/repository"
	"fmt"
)

type CollectionService struct {
	repo *repository.CollectionRepo
}

func NewCollectionService() *CollectionService {
	return &CollectionService{repo: &repository.CollectionRepo{}}
}

func (s *CollectionService) Collect(userID, postID uint64) error {
	if s.repo.Exists(userID, postID) {
		return fmt.Errorf("already collected")
	}

	collection := &model.Collection{UserID: userID, PostID: postID}
	return s.repo.Create(collection)
}
