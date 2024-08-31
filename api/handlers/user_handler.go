package api

import (
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

func (h *UserHanlder) HandlePostUser(c *fiber.Ctx) error {
	var params types.CreateUserParams
	if err := c.BodyParser(&params); err != nil {
		return err
	}

	if errors := params.Validate(); len(errors) > 0 {
		return c.JSON(errors)
	}

	user, err := types.NewUserFromParmas(params)
	if err != nil {
		return nil
	}

	insertedUser, err := h.userStore.InsertUser(c.Context(), user)
	if err != nil {
		return err
	}
	return c.JSON(insertedUser)
}

func (h *UserHanlder) HandleGetUser(c *fiber.Ctx) error {
	var (
		id = c.Params("id")
	)
	user, err := h.userStore.GetUserByID(c.Context(), id)
	if err != nil {
		return err
	}
	return c.JSON(user)
}

func (h *UserHanlder) HandleGetUsers(c *fiber.Ctx) error {

	users, err := h.userStore.GetUsers(c.Context())
	if err != nil {
		return err
	}
	return c.JSON(users)
}
