package message

import (
	"context"
	"errors"
	"regexp"
	"strings"

	"github.com/mhdiiilham/segrato/user"
)

type service struct {
	messageRepository Repository
	userRepository    user.Repository
}

func NewService(messageRepository Repository, userRepository user.Repository) Service {
	return &service{
		messageRepository: messageRepository,
		userRepository:    userRepository,
	}
}

func (s service) PostMessage(ctx context.Context, msg Message) (ID string, err error) {
	userBlockedWords, err := s.userRepository.GetUserBlockedWords(ctx, msg.UserID)
	if err != nil {
		return
	}

	words := strings.Join(userBlockedWords, "|")
	re := regexp.MustCompile(`(?i)` + words)
	if blocked := re.MatchString(msg.Message); blocked {
		err = errors.New("message contained banned words from user")
		return
	}
	return s.messageRepository.Create(ctx, msg)
}

func (s service) GetUserMessages(ctx context.Context, userID string) (messages []Message, err error) {
	return s.messageRepository.FindByUserID(ctx, userID)
}

func (s service) GetMessage(ctx context.Context, msgID string) (msg Message, err error) {
	return s.messageRepository.FindOne(ctx, msgID)
}

func (s service) ReplyToAMessage(ctx context.Context, msgID string, replyMessage RepliedMessage) (msg Message, err error) {
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
