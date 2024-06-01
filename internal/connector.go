package internal

import (
	"mnc-finance/config"
	"mnc-finance/internal/usecase"
	"sync"
)

var oneUc sync.Once
var uc usecase.Usecase

func RegisterUc(cfg *config.Config) usecase.Usecase {
	oneUc.Do(func() {
		uc = usecase.ProvideUsecase(
			cfg,
		)
	})

	return uc
}
