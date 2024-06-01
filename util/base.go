package util

import (
	"mnc-finance/internal/contract/entity"
)

var (
	Items = []entity.Language{}
)

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
