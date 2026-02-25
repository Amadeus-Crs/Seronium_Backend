package service

import (
	"Seronium/internal/model"
	"Seronium/internal/repository"
	"encoding/json"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type PostService struct {
	repo *repository.PostRepo
}

func NewPostService() *PostService {
	return &PostService{repo: &repository.PostRepo{}}
}

func (s *PostService) Create(userID uint64, title, content, typ string) (uint64, error) {
	post := &model.Post{UserID: userID, Title: title, Content: content, Type: typ}
	if err := s.repo.Create(post); err != nil {
		return 0, err
	}
	repository.RDB.ZAdd(repository.CTX, "post:hot", redis.Z{Score: 0, Member: post.ID})
	return post.ID, nil
}

func (s *PostService) Get(id uint64) (*model.Post, error) {
	key := fmt.Sprintf("post:detail:%d", id)
	data, err := repository.RDB.Get(repository.CTX, key).Bytes()
	if err == nil {
		var post model.Post
		if err := json.Unmarshal(data, &post); err == nil {
			return &post, nil
		}
	}

	post, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	jsonData, _ := json.Marshal(post)
	repository.RDB.Set(repository.CTX, key, jsonData, 15*time.Minute+time.Duration(rand.Intn(300))*time.Second)

	return post, nil
}

func (s *PostService) Update(userID, id uint64, title, content string) error {
	post, err := s.repo.FindByID(id)
	if err != nil || post.UserID != userID {
		return fmt.Errorf("post not found or permission denied")
	}
	post.Title = title
	post.Content = content
	if err := s.repo.Update(post); err != nil {
		return err
	}
	key := fmt.Sprintf("post:detail:%d", id)
	repository.RDB.Del(repository.CTX, key)
	return nil
}

func (s *PostService) Delete(userID, id uint64) error {
	post, err := s.repo.FindByID(id)
	if err != nil || post.UserID != userID {
		return fmt.Errorf("post not found or permission denied")
	}
	if err := s.repo.Delete(id); err != nil {
		return err
	}
	repository.RDB.ZRem(repository.CTX, "post:hot", id)
	key := fmt.Sprintf("post:detail:%d", id)
	repository.RDB.Del(repository.CTX, key)
	return nil
}

func (s *PostService) List(offset, limit int, sort string) ([]model.Post, error) {
	if sort == "hot" {
		idsStr, err := repository.RDB.ZRangeArgs(repository.CTX, redis.ZRangeArgs{
			Key:   "post:hot",
			Start: 0,
			Stop:  int64(limit - 1),
		}).Result()
		if err != nil {
			zap.L().Error("redis zrange failed", zap.Error(err))
			return nil, err
		}
		var ids []uint64
		for _, idStr := range idsStr {
			id, _ := strconv.ParseUint(idStr, 10, 64)
			ids = append(ids, id)
		}
		var posts []model.Post
		repository.DB.Where("id IN ?", ids).Find(&posts)
		return posts, nil
	}
	return s.repo.List(offset, limit)
}

func (s *PostService) IncView(id uint64) {
	repository.RDB.ZIncrBy(repository.CTX, "post:hot", 1, strconv.FormatUint(id, 10))
	repository.DB.Model(&model.Post{}).Where("id = ?", id).Update("view_count", gorm.Expr("view_count + ?", 1))
}
