package user

import "context"

type Repository interface {
	Create(context.Context, User) (User, error)
	FindOne(ctx context.Context, username string) (user User, err error)
	FindByID(ctx context.Context, id string) (user User, err error)
	GetUserBlockedWords(ctx context.Context, userID string) (blockedWords []string, err error)
	CheckUniqueness(ctx context.Context, username string) (unique bool)
}

type Service interface {
	RegisterUser(ctx context.Context, username, plainPassword string, blockWords []string) (user User, accessToken string, err error)
	GetUser(ctx context.Context, username string) (user User, err error)
}
