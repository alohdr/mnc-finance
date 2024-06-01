package usecase

import "github.com/gofiber/fiber/v2"

func (u *UsecaseImpl) HealthCheck(c *fiber.Ctx) error {
	err := u.Repo.Ping()
	if err != nil {
		return err
	}

	return nil
}
