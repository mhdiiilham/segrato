package service

import (
	"context"

	"github.com/mhdiiilham/segrato/message"
	"github.com/mhdiiilham/segrato/user"
)

type messageService struct {
	messageRepository message.Repository
	userRepository    user.Repository
}

func NewMessageService(messageRepository message.Repository, userRepository user.Repository) message.Service {
	return &messageService{
		messageRepository: messageRepository,
		userRepository:    userRepository,
	}
}

func (s messageService) PostMessage(ctx context.Context, msg message.Message) (ID string, err error) {
	return s.messageRepository.Create(ctx, msg)
}

func (s messageService) GetUserMessages(ctx context.Context, userID string) (messages []message.Message, err error) {
	return s.messageRepository.FindByUserID(ctx, userID)
}

func (s messageService) GetMessage(ctx context.Context, msgID string) (msg message.Message, err error) {
	return s.messageRepository.FindOne(ctx, msgID)
}

func (s messageService) ReplyToAMessage(ctx context.Context, msgID string, replyMessage message.RepliedMessage) (msg message.Message, err error) {
	msg, err = s.GetMessage(ctx, msgID)
	if err != nil {
		return
	}

	msg.Replied = append(msg.Replied, replyMessage)
	if err = s.messageRepository.UpdateOne(ctx, msg); err != nil {
		return
	}
	return
}