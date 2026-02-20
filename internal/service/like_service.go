package service

import (
	"Seronium/internal/model"
	"Seronium/internal/repository"
	"fmt"
	"strconv"
	"time"
)

type LikeService struct {
	repo *repository.LikeRepo
}

func NewLikeService() *LikeService {
	return &LikeService{repo: &repository.LikeRepo{}}
}

func (s *LikeService) Like(userID uint64, targetType string, targetID uint64) error {
	key := fmt.Sprintf("rate:like:%d", userID)
	count, err := repository.RDB.Incr(repository.CTX, key).Result()
	if err != nil {
		return err
	}
	if count == 1 {
		repository.RDB.Expire(repository.CTX, key, time.Minute)
	}
	if count > 5 {
		return fmt.Errorf("rate limit exceeded")
	}

	if s.repo.Exists(userID, targetType, targetID) {
		return fmt.Errorf("already liked")
	}

	like := &model.Like{UserID: userID, TargetType: targetType, TargetID: targetID}
	if err := s.repo.Create(like); err != nil {
		return err
	}

	if targetType == "post" {
		repository.RDB.ZIncrBy(repository.CTX, "post:hot", 5, strconv.FormatUint(targetID, 10))
	}
	return nil
}
