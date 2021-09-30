package message

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Message struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	SenderID   string             `bson:"senderId" json:"senderId"`
	SenderName string             `bson:"senderName" json:"senderName"`
	UserID     string             `bson:"userId" json:"userId"`
	Message    string             `bson:"message" json:"message"`
	Replied    []RepliedMessage   `bson:"replied" json:"replied"`
	Show       bool               `bson:"show" json:"show"`
	CreatedAt  time.Time          `bson:"createdAt" json:"createdAt"`
}

type RepliedMessage struct {
	SenderID   string `bson:"senderId" json:"senderId"`
	SenderName string `bson:"senderName" json:"senderName"`
	Message    string `bson:"message" json:"message"`
}

type PostMessageRequest struct {
	Message    string `json:"message"`
	UserID     string `json:"userId"`
	SenderID   string `json:"senderId"`
	SenderName string `json:"senderName"`
}

type GetUserMessagesResponse struct {
	Code          int       `json:"code"`
	UserID        string    `json:"userId"`
	Messages      []Message `json:"messages"`
	TotalMessages int       `json:"totalMessage"`
}

const (
	BannedWord string = "message contained banned words from user"
)
