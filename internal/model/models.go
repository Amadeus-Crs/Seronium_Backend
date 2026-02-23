package model

import "time"

type User struct {
	ID           uint64    `gorm:"primaryKey;autoIncrement"`
	Username     string    `gorm:"type:varchar(50);unique;not null"`
	PasswordHash string    `gorm:"type:varchar(255);not null"`
	Salt         string    `gorm:"type:varchar(64);not null"`
	AvatarURL    string    `gorm:"type:varchar(255)"`
	Bio          string    `gorm:"type:text"`
	Role         string    `gorm:"type:enum('normal','vip','admin','banned');default:'normal'"`
	CreatedAt    time.Time `gorm:"autoCreateTime"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime"`
}

type Post struct {
	ID        uint64 `gorm:"primaryKey;autoIncrement"`
	UserID    uint64
	Title     string `gorm:"not null"`
	Content   string `gorm:"not null"`
	Type      string `gorm:"type:enum('question','article');not null"` // question/article
	Status    string `gorm:"default:'published'"`
	ViewCount uint64 `gorm:"default:0"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Comment struct {
	ID         uint64 `gorm:"primaryKey;autoIncrement"`
	UserID     uint64
	TargetType string `gorm:"not null"` // post/answer
	TargetID   uint64 `gorm:"not null"`
	Content    string `gorm:"not null"`
	CreatedAt  time.Time
}

type Like struct {
	UserID     uint64 `gorm:"primaryKey"`
	TargetType string `gorm:"primaryKey"`
	TargetID   uint64 `gorm:"primaryKey"`
	CreatedAt  time.Time
}

type Collection struct {
	UserID    uint64 `gorm:"primaryKey"`
	PostID    uint64 `gorm:"primaryKey"`
	CreatedAt time.Time
}
