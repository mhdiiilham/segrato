package auth

import (
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/mhdiiilham/segrato/internal/auth/model"
	"go.mongodb.org/mongo-driver/mongo"
)

func (s *Server) GetUser(c *fiber.Ctx) error {
	userID := c.Params("userid")
	user, err := s.userService.GetUser(c.Context(), userID)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return c.Status(http.StatusNotFound).JSON(model.Error{
				Code:    http.StatusNotFound,
				Message: "user not found",
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

	return c.Status(http.StatusOK).JSON(user)
}
