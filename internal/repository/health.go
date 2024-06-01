package repository

import "github.com/gofiber/fiber/v2/log"

func (r *RepositoryImpl) Ping() error {

	sqlDB, err := r.db.DB()
	if err != nil {
		return err
	}

	err = sqlDB.Ping()
	if err != nil {
		log.Errorf("error on db connection : ", err)
		return err
	}

	return nil
}
