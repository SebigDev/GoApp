package handlers

import (
	"CrashCourse/GoApp/src/modules/dto"
	"CrashCourse/GoApp/src/modules/services"

	"time"

	"github.com/gofiber/fiber/v2"
)

type IAuthHandler interface {
	Authenticate(ctx *fiber.Ctx) error
}

type authHandler struct {
	UserService services.IUserService
}

func NewAuthHandler(userService services.IUserService) IAuthHandler {
	return &authHandler{
		UserService: userService,
	}
}

func (s authHandler) Authenticate(ctx *fiber.Ctx) error {
	authRequest := new(dto.LoginRequest)

	if err := ctx.BodyParser(authRequest); err != nil {
		return ctx.Status(500).JSON("Error parsing request")
	}

	tokenResponse, err := s.UserService.LoginPerson(*authRequest)

	if err != nil {
		return ctx.Status(401).JSON(err.Error())
	}

	cookie := fiber.Cookie{
		Name:     "refreshToken",
		Value:    tokenResponse.RefreshToken,
		Expires:  time.Now().UTC().Add(time.Hour * 24),
		HTTPOnly: true,
		Secure:   true,
	}

	ctx.Cookie(&cookie)
	return ctx.JSON(fiber.Map{"token": tokenResponse.Token})
}