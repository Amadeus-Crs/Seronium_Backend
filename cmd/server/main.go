package main

import (
	"Seronium/internal/config"
	"Seronium/internal/model"
	"Seronium/internal/repository"
	"log"
)

func main() {
	if err := config.Init(); err != nil {
		log.Fatalf("failed to initialize config: %v", err)
	}
	if err := repository.InitDB(); err != nil {
		log.Fatalf("failed to initialize database: %v", err)
	}
	member1 := &repository.UserRepo{}

	if err := member1.Create(&model.User{
		Username:     "testuser",
		PasswordHash: "hashedpassword",
		Salt:         "randomsalt",
		AvatarURL:    "http://example.com/avatar.jpg",
		Bio:          "This is a test user.",
		Role:         "normal",
	}); err != nil {
		log.Fatalf("failed to create user: %v", err)
	}
	log.Println("User created successfully")
}
