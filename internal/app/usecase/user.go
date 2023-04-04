package usecase

import (
	"context"
	"graphql-golang/internal/app/entity"
	"graphql-golang/internal/app/repository"
)

type UserUseCase interface {
	CreateUser(context.Context, entity.CreateUserRequest) (*entity.CreateUserResponse, error)
}

type userUseCase struct {
	userRepository repository.UserRepository
}

func NewUserUseCase(userRepository repository.UserRepository) UserUseCase {
	return &userUseCase{
		userRepository: userRepository,
	}
}

func (u userUseCase) CreateUser(ctx context.Context, request entity.CreateUserRequest) (*entity.CreateUserResponse, error) {
	//TODO implement me
	panic("implement me")
}
