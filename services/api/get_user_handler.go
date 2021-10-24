package api

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (s *Server) GetUser(c *fiber.Ctx) error {
	userID := c.Params("userid")
	user, err := s.userService.GetUser(c.Context(), userID)
	if err != nil {
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
