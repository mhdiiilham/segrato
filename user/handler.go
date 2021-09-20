package user

import (
	"context"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/mhdiiilham/segrato/pkg/token"
)

type RegisterUserResponse struct {
	ID          string `json:"id"`
	Username    string `json:"username"`
	AccessToken string `json:"accessToken"`
}

type RegisterUserPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Handler struct {
	userRepository Repository
	tokenService   token.Service
}

func NewHandler(userRepository Repository, tokenService token.Service) Handler {
	return Handler{userRepository: userRepository, tokenService: tokenService}
}

func (h Handler) RegisterUser(c *fiber.Ctx) error {
	var accessToken string
	var payload RegisterUserPayload
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
		}{
			Code:    http.StatusInternalServerError,
			Message: "Internal Server Error",
		})
	}

	if len(payload.Password) < 8 || len(payload.Username) < 5 {
		return c.Status(http.StatusBadRequest).JSON(struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
		}{
			Code:    http.StatusBadRequest,
			Message: "Username or/and password Not Valid",
		})
	}

	user, err := h.userRepository.Create(context.Background(), payload.Username, payload.Password)
	if err != nil {
		if err.Error() == "username already taken" {
			return c.Status(http.StatusBadRequest).JSON(struct {
				Code    int    `json:"code"`
				Message string `json:"message"`
			}{
				Code:    http.StatusBadRequest,
				Message: "Username already taken",
			})
		}

		return c.Status(http.StatusInternalServerError).JSON(struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
		}{
			Code:    http.StatusInternalServerError,
			Message: "Internal Server Error",
		})
	}

	accessToken, err = h.tokenService.SignPayload(token.TokenPayload{
		ID:        user.ID.Hex(),
		Username:  user.Username,
		IsPremium: user.IsPremium,
	})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
		}{
			Code:    http.StatusInternalServerError,
			Message: "Internal Server Error",
		})
	}

	resp := RegisterUserResponse{
		ID:          user.ID.String(),
		Username:    payload.Username,
		AccessToken: accessToken,
	}
	return c.Status(http.StatusCreated).JSON(resp)
}
