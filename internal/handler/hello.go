package handler

import (
	"github.com/gofiber/fiber/v2"
	"mnc-finance/util"
)

func (h *Handler) Hello(ctx *fiber.Ctx) error {
	return ctx.JSON(util.Response{
		Success: true,
		Message: "Hello Go developers",
		Data:    nil,
	})
}
