package handlers

import (
	"context"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/useresd/golang-fiber-starter/internal/database"
	"github.com/useresd/golang-fiber-starter/internal/models"
	"github.com/useresd/golang-fiber-starter/internal/services"
)

type UserHandler struct {
	userService *services.UserService
}

func NewUserHandler(userService *services.UserService) *UserHandler {
	return &UserHandler{userService}
}

func (h *UserHandler) HandleGetUsers(c *fiber.Ctx) error {

	var users []*models.User

	err := database.BeginTx(c.UserContext(), func(ctx context.Context) error {

		var err error

		if users, err = h.userService.FindMany(ctx); err != nil {
			return err
		}

		return nil

	})

	if err != nil {
		return err
	}

	return c.JSON(map[string]any{
		"data": users,
	})
}

func (h *UserHandler) HandlePostUser(c *fiber.Ctx) error {

	user := &models.User{}

	err := c.BodyParser(user)

	if err != nil {
		return err
	}

	err = database.BeginTx(c.UserContext(), func(ctx context.Context) error {

		return h.userService.Store(ctx, user)

	})

	if err != nil {
		return err
	}

	return c.SendStatus(http.StatusNoContent)

}
