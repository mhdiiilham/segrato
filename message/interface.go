package message

import (
	"context"
)

type Repository interface {
	Create(ctx context.Context, msg Message) (ID string, err error)
	FindOne(ctx context.Context, id string) (msg Message, err error)
	FindByUserID(ctx context.Context, userID string) (messages []Message, err error)
	UpdateOne(ctx context.Context, msg Message) (err error)
}

type Service interface {
	PostMessage(ctx context.Context, msg Message) (ID string, err error)
	GetUserMessages(ctx context.Context, userID string) (messages []Message, err error)
	GetMessage(ctx context.Context, msgID string) (msg Message, err error)
	ReplyToAMessage(ctx context.Context, msgID string, replyMessage RepliedMessage) (msg Message, err error)
}
