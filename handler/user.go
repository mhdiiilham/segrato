package handler

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/mhdiiilham/segrato/user"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{
		userService: userService,
	}
}

func (h userHandler) RegisterUser(ctx *fiber.Ctx) error {
	var accessToken string
	var payload user.RegisterUserPayload
	if err := ctx.BodyParser(&payload); err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(NewInternalServerError())
	}

	if len(payload.Password) < 8 || len(payload.Username) < 5 {
		return ctx.Status(http.StatusBadRequest).JSON(ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Username or/and Password Not Valid",
		})
	}

	u, accessToken, err := h.userService.RegisterUser(ctx.Context(), payload.Username, payload.Password, payload.BlockedWords)
	if err != nil {
		if err.Error() == "username already taken" {
			return ctx.Status(http.StatusBadRequest).JSON(ErrorResponse{
				Code:    http.StatusBadRequest,
				Message: err.Error(),
			})
		}

		return ctx.Status(http.StatusInternalServerError).JSON(NewInternalServerError())
	}

	return ctx.Status(http.StatusCreated).JSON(user.RegisterUserResponse{
		ID:          u.ID.Hex(),
		User:        u,
		AccessToken: accessToken,
	})
}
