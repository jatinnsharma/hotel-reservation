package api

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/jatinnsharma/hotel-reservation/db"
	"github.com/jatinnsharma/hotel-reservation/types"
)

type UserHanlder struct {
	userStore db.UserStore
}

func NewUserHandler(userStore db.UserStore) *UserHanlder {
	return &UserHanlder{
		userStore: userStore,
	}
}

func (h *UserHanlder) HandleGetUsers(c *fiber.Ctx) error {
	u := types.User{
		FirstName: "James",
		LastName:  "At the watercooler",
	}
	return c.JSON(u)
}

func (h *UserHanlder) HandleGetUser(c *fiber.Ctx) error {
	var (
		id  = c.Params("id")
		ctx = context.Background()
	)
	user, err := h.userStore.GetUserByID(ctx, id)
	if err != nil {
		return err
	}
	return c.JSON(user)
}
