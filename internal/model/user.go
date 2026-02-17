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
