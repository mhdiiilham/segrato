package auth

import (
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/mhdiiilham/segrato/internal/apiresponse"
	"github.com/mhdiiilham/segrato/internal/auth/model/user"
	"go.mongodb.org/mongo-driver/mongo"
)

func (s *Server) Login(c *fiber.Ctx) error {
	var requestBody user.LoginRequest

	if err := c.BodyParser(&requestBody); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(apiresponse.Error{
			Code:    http.StatusInternalServerError,
			Message: "internal server error",
		})
	}

	if len(requestBody.Username) < 5 || len(requestBody.Password) < 8 {
		return c.Status(http.StatusBadRequest).JSON(apiresponse.Error{
			Code:    http.StatusBadRequest,
			Message: "invalid username and/or password",
		})
	}

	userData, accessToken, err := s.userService.Login(c.Context(), requestBody.Username, requestBody.Password)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return c.Status(http.StatusBadRequest).JSON(apiresponse.Error{
				Code:    http.StatusBadRequest,
				Message: user.ErrInvalidUsernamePassword.Error(),
			})
		}

		if errors.Is(err, user.ErrInvalidUsernamePassword) {
			return c.Status(http.StatusBadRequest).JSON(apiresponse.Error{
				Code:    http.StatusBadRequest,
				Message: err.Error(),
			})
		}

		return c.Status(http.StatusInternalServerError).JSON(apiresponse.Error{
			Code:    http.StatusInternalServerError,
			Message: "internal server error",
		})
	}

	return c.Status(http.StatusOK).JSON(user.LoginResponse{
		ID:          userData.ID.Hex(),
		AccessToken: accessToken,
		User:        userData,
	})

}
