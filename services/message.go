package service

import (
	"context"
	"errors"
	"strings"

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
	userBlockedWords, err := s.userRepository.GetUserBlockedWords(ctx, msg.UserID)
	if err != nil {
		return
	}

	for _, word := range userBlockedWords {
		// TODO: Need to change this using regex
		text := strings.ToLower(msg.Message)
		if blocked := strings.Contains(text, strings.ToLower(word)); blocked {
			err = errors.New("message contained banned words from user")
			return
		}
	}
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
