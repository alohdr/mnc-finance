package handler

import (
	"github.com/gofiber/fiber/v2"
	"mnc-finance/internal/usecase"
)

type Handler struct {
	Usecase usecase.Usecase
}

func NewHandler(Usecase usecase.Usecase) Handler {
	return Handler{Usecase: Usecase}
}

func (h *Handler) Mount(router fiber.Router) {
	router.Get("/health", h.HealthCheck)

	router.Get("/", h.Hello)

	router.Get("/palindrome-check", h.Palindrome)

	router.Get("/language", h.Language)
	router.Post("/language", h.AddLanguage)
	router.Get("/language/:id", h.GetDetailLanguage)
	router.Get("/languages", h.GetListLanguage)
	router.Patch("/language/:id", h.UpdateLanguage)
	router.Delete("/language/:id", h.DeleteLanguage)

	//masterData := router.Group("/master-data")
	//
	//masterData.Post("/pol", h.CreatePol)
}
