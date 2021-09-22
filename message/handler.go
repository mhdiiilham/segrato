package message

import (
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

type handler struct {
	messageRepository Repository
}

func NewHandler(messageRepository Repository) *handler {
	return &handler{
		messageRepository: messageRepository,
	}
}

func (h handler) PostMessage(ctx *fiber.Ctx) error {
	var payload PostMessageRequest
	if err := ctx.BodyParser(&payload); err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
		}{
			Code:    http.StatusInternalServerError,
			Message: "Internal Server Error",
		})
	}

	msg := Message{
		SenderID:   payload.SenderID,
		SenderName: payload.SenderName,
		UserID:     payload.UserID,
		Message:    payload.Message,
		Replied:    []RepliedMessage{},
		CreatedAt:  time.Now(),
	}

	_, err := h.messageRepository.Create(ctx.Context(), msg)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
		}{
			Code:    http.StatusInternalServerError,
			Message: "Internal Server Error",
		})
	}

	return ctx.Status(http.StatusCreated).JSON(msg)
}

func (h handler) GetUserMessages(ctx *fiber.Ctx) error {
	userID := ctx.Params("id")
	messages, err := h.messageRepository.FindByUserID(ctx.Context(), userID)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
		}{
			Code:    http.StatusInternalServerError,
			Message: "Internal Server Error",
		})
	}

	return ctx.Status(http.StatusOK).JSON(GetUserMessagesResponse{
		Code:     http.StatusOK,
		UserID:   userID,
		Messages: messages,
	})
}

func (h handler) RepliedMessage(ctx *fiber.Ctx) error {
	msgID := ctx.Params("id")
	var payload RepliedMessage
	if err := ctx.BodyParser(&payload); err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
		}{
			Code:    http.StatusInternalServerError,
			Message: "Internal Server Error",
		})
	}

	msg, err := h.messageRepository.FindOne(ctx.Context(), msgID)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return ctx.Status(http.StatusBadRequest).JSON(struct {
				Code    int    `json:"code"`
				Message string `json:"message"`
			}{
				Code:    http.StatusBadRequest,
				Message: "invalid message id",
			})
		}

		return ctx.Status(http.StatusInternalServerError).JSON(struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
		}{
			Code:    http.StatusInternalServerError,
			Message: "Internal Server Error",
		})
	}

	msg.Replied = append(msg.Replied, RepliedMessage{Message: payload.Message, SenderID: payload.SenderID, SenderName: payload.SenderName})
	if err := h.messageRepository.UpdateOne(ctx.Context(), msg); err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
		}{
			Code:    http.StatusInternalServerError,
			Message: "Internal Server Error",
		})
	}

	return ctx.Status(http.StatusOK).JSON(msg)
}
