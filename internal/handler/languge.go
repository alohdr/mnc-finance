package handler

import (
	"github.com/gofiber/fiber/v2"
	"mnc-finance/internal/contract/model"
	"mnc-finance/util"
	"net/http"
)

func (h *Handler) Language(ctx *fiber.Ctx) error {
	param := model.Language{}

	err := ctx.BodyParser(&param)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(util.Response{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return ctx.JSON(util.Response{
		Success: true,
		Message: util.Success,
		Data:    param,
	})
}

func (h *Handler) AddLanguage(ctx *fiber.Ctx) error {
	param := model.Language{}

	err := ctx.BodyParser(&param)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(util.Response{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	statusCode, status, err := h.Usecase.AddLanguage(param)
	if err != nil {
		return ctx.Status(statusCode).JSON(util.Response{
			Success: status,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return ctx.JSON(util.Response{
		Success: status,
		Message: util.Success,
		Data:    nil,
	})
}

func (h *Handler) UpdateLanguage(ctx *fiber.Ctx) error {
	param := model.Language{}

	err := ctx.BodyParser(&param)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(util.Response{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	statusCode, status, err := h.Usecase.UpdateLanguage(param, ctx.Params("id"))
	if err != nil {
		return ctx.Status(statusCode).JSON(util.Response{
			Success: status,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return ctx.JSON(util.Response{
		Success: status,
		Message: util.Success,
		Data:    nil,
	})
}

func (h *Handler) GetDetailLanguage(ctx *fiber.Ctx) error {
	resp, statusCode, status, err := h.Usecase.GetDetailLanguage(ctx.Params("id"))
	if err != nil {
		return ctx.Status(statusCode).JSON(util.Response{
			Success: status,
			Message: err.Error(),
			Data:    resp,
		})
	}

	return ctx.JSON(util.Response{
		Success: status,
		Message: util.Success,
		Data:    resp,
	})
}

func (h *Handler) DeleteLanguage(ctx *fiber.Ctx) error {
	statusCode, status, err := h.Usecase.DeleteLanguage(ctx.Params("id"))
	if err != nil {
		return ctx.Status(statusCode).JSON(util.Response{
			Success: status,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return ctx.JSON(util.Response{
		Success: status,
		Message: util.Success,
		Data:    nil,
	})
}

func (h *Handler) GetListLanguage(ctx *fiber.Ctx) error {
	resp, statusCode, status, err := h.Usecase.GetListLanguage()
	if err != nil {
		return ctx.Status(statusCode).JSON(util.Response{
			Success: status,
			Message: err.Error(),
			Data:    resp,
		})
	}

	return ctx.JSON(util.Response{
		Success: status,
		Message: util.Success,
		Data:    resp,
	})
}
