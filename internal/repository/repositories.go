package repository

import (
	"gorm.io/gorm"
)

type RepositoryImpl struct {
	db *gorm.DB
}

type Repository interface {
	Ping() error
}

func ProvideRepository(db *gorm.DB) Repository {
	return &RepositoryImpl{
		db: db,
	}
}
