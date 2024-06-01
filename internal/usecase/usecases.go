package usecase

import (
	"github.com/gofiber/fiber/v2"
	"mnc-finance/config"
	"mnc-finance/internal/contract/model"
	"mnc-finance/internal/repository"
)

type Usecase interface {
	HealthCheck(c *fiber.Ctx) error
	Palindrome(param model.PalindromeRequest) (string, bool)
	AddLanguage(param model.Language) (int, bool, error)
	UpdateLanguage(param model.Language, id string) (int, bool, error)
	GetDetailLanguage(id string) (*model.Language, int, bool, error)
	GetListLanguage() ([]model.Language, int, bool, error)
	DeleteLanguage(id string) (int, bool, error)
}

type UsecaseImpl struct {
	Cfg  *config.Config
	Repo repository.Repository
}

func ProvideUsecase(cfg *config.Config) *UsecaseImpl {
	return &UsecaseImpl{
		Cfg: cfg,
	}
}
