package postgres

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Host     string
	Port     string
	Name     string
	User     string
	Password string
}

func NewPostgresDB(cfg *Config) (*gorm.DB, error) {
	dsn := fmt.Sprint(
		"postgres://" + cfg.User +
			":" + cfg.Password +
			"@" + cfg.Host +
			":" + cfg.Port +
			"/" + cfg.Name)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
