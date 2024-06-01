package handler

import (
	"github.com/gofiber/fiber/v2"
	"mnc-finance/internal/contract/model"
	"mnc-finance/util"
	"net/http"
)

func (h *Handler) Palindrome(ctx *fiber.Ctx) error {
	params := model.PalindromeRequest{}

	err := ctx.QueryParser(&params)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(util.Response{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	msg, status := h.Usecase.Palindrome(params)
	if !status {
		return ctx.Status(http.StatusBadRequest).JSON(util.Response{
			Success: status,
			Message: fiber.ErrBadRequest.Message,
			Data:    msg,
		})
	}

	return ctx.JSON(util.Response{
		Success: status,
		Message: util.Success,
		Data:    msg,
	})
}
