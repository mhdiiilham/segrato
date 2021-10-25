package user

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/mhdiiilham/segrato/pkg/token"
)

type Handler struct {
	userRepository Repository
	tokenService   token.Service
}

func NewHandler(userRepository Repository, tokenService token.Service) *Handler {
	return &Handler{
		userRepository: userRepository,
		tokenService:   tokenService,
	}
}

func (h Handler) GetUser(ctx *fiber.Ctx) error {
	userID := ctx.Params("userid")
	user, err := h.userRepository.FindByID(ctx.Context(), userID)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
		}{
			Code:    http.StatusInternalServerError,
			Message: "Internal Server Error " + err.Error(),
		})
	}

	return ctx.Status(http.StatusOK).JSON(user)
}
