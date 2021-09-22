package user

import (
	"context"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/mhdiiilham/segrato/pkg/token"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

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
		User:        user,
		AccessToken: accessToken,
	}
	return c.Status(http.StatusCreated).JSON(resp)
}

func (h Handler) Login(ctx *fiber.Ctx) error {
	var payload RegisterUserPayload

	if err := ctx.BodyParser(&payload); err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
		}{
			Code:    http.StatusInternalServerError,
			Message: "Internal Server Error",
		})
	}

	if len(payload.Password) < 8 || len(payload.Username) < 5 {
		return ctx.Status(http.StatusBadRequest).JSON(struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
		}{
			Code:    http.StatusBadRequest,
			Message: "Username or/and password Not Valid",
		})
	}

	user, err := h.userRepository.FindOne(ctx.Context(), payload.Username)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return ctx.Status(http.StatusBadRequest).JSON(struct {
				Code    int    `json:"code"`
				Message string `json:"message"`
			}{
				Code:    http.StatusBadRequest,
				Message: "Wrong Username or/and Password",
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

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password)); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
		}{
			Code:    http.StatusBadRequest,
			Message: "Wrong Username or/and Password",
		})
	}

	accessToken, accessTokenErr := h.tokenService.SignPayload(token.TokenPayload{ID: user.ID.Hex(), Username: user.Username, IsPremium: user.IsPremium})
	if accessTokenErr != nil {
		return ctx.Status(http.StatusBadRequest).JSON(struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
		}{
			Code:    http.StatusBadRequest,
			Message: "Wrong Username or/and Password",
		})
	}

	return ctx.Status(http.StatusOK).JSON(RegisterUserResponse{
		ID:          user.ID.Hex(),
		User:        user,
		AccessToken: accessToken,
	})
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
