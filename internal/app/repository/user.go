package repository

import (
	"context"
	"graphql-golang/internal/app/entity"
	// "gorm.io/gorm"
)

type UserRepository interface {
	Create(context.Context, entity.CreateUserRequest) (*entity.CreateUserResponse, error)
}

// type userRepository struct {
// 	db *gorm.DB
// }

// func NewUserRepository(db *gorm.DB) UserRepository {
// 	return &userRepository{
// 		db: db,
// 	}
// }

// func (u userRepository) Create(ctx context.Context, request entity.CreateUserRequest) (*entity.CreateUserResponse, error) {
// 	//TODO implement me
// 	panic("implement me")
// }
