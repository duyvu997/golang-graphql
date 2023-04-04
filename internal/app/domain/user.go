package domain

import (
	"context"
	"graphql-golang/internal/app/entity"
)

func CreateUser(ctx context.Context, input entity.CreateUserRequest) (entity.CreateUserResponse, error) {
	return entity.CreateUserResponse{}, nil
}
