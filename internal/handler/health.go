package handler

import (
	"github.com/gofiber/fiber/v2"
	"mnc-finance/util"
	"net/http"
)

func (h *Handler) HealthCheck(ctx *fiber.Ctx) error {
	err := h.Usecase.HealthCheck(ctx)
	if err != nil {
		return fiber.NewError(http.StatusBadRequest, err.Error())
	}

	return ctx.JSON(util.Response{
		Success: true,
		Message: `Its healthy`,
		Data:    nil,
	})
}
