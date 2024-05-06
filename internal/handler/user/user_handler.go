package handler

import (
	model "king/internal/domain/model/user"
	usecase "king/internal/usecase/user"
	"log"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	Userusecase *usecase.UserUsecase
}

func NewUserHandler(userhandler *usecase.UserUsecase) *UserHandler {
	return &UserHandler{Userusecase: userhandler}
}

func (h *UserHandler) RegisterUser(c *fiber.Ctx) error {
	var userReq model.User
	if err := c.BodyParser(&userReq); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
	}
	
	if err := h.Userusecase.RegisterUser(userReq.Username, userReq.Password); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(userReq)

}

func(h *UserHandler) GetUserName(c *fiber.Ctx) error {

	username := c.Params("username")

	user, err := h.Userusecase.GetUserByUsername(username)
	if err != nil {
        // Check if the error indicates that the user was not found
        if err == fiber.ErrNotFound {
            // Return a 404 Not Found response
            return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
                "error": "User not found",
            })
        }
        // Log other errors
        log.Println("Error:", err)
        // Return a generic error response
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Internal Server Error",
        })
    }
 
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "successfully", "info": user});

}