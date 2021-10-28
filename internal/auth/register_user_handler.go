package auth

import (
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/mhdiiilham/segrato/internal/auth/model"
	"github.com/mhdiiilham/segrato/internal/auth/model/user"
	"github.com/sirupsen/logrus"
)

func (s *Server) RegisterUser(c *fiber.Ctx) error {
	var payload model.UserRegiserPayload

	logrus.Info("creating new user")
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(http.StatusBadRequest).JSON(model.Error{
			Code:    http.StatusBadRequest,
			Message: "username and/or password invalid",
		})
	}

	if len(payload.Password) < 8 || len(payload.Username) < 5 || payload.Email == "" {
		return c.Status(http.StatusBadRequest).JSON(model.Error{
			Code:    http.StatusBadRequest,
			Message: "username, email and/or password invalid",
		})
	}

	newUser, accessToken, err := s.userService.RegisterUser(c.Context(), payload.Username, payload.Email, payload.Password)
	if err != nil {
		if errors.Is(err, user.ErrUsernameEmailRegistered) {
			return c.Status(http.StatusBadRequest).JSON(model.Error{
				Code:    http.StatusBadRequest,
				Message: err.Error(),
			})
		}

		logrus.Errorf("error trying to create new user: %v", err)
		return c.Status(http.StatusInternalServerError).JSON(model.Error{
			Code:    http.StatusInternalServerError,
			Message: "internal server error",
		})
	}

	logrus.Infof("new user created with id %s", newUser.ID.Hex())
	return c.Status(http.StatusOK).JSON(model.ResgisterUserResponse{
		ID:          newUser.ID.Hex(),
		User:        newUser,
		AccessToken: accessToken,
	})
}
