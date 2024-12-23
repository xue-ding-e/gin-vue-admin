package system

import (
	"github.com/gofiber/fiber/v2"
)

type BaseRouter struct{}

func (s *BaseRouter) InitBaseRouter(router fiber.Router) fiber.Router {
	baseRouter := router.Group("base")
	{
		baseRouter.Post("/login", baseApi.Login)
		baseRouter.Post("/captcha", baseApi.Captcha)
	}
	return baseRouter
}
