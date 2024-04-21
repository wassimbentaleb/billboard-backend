package database

import (
	"billboard/entities"
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Postgres struct {
	DB *gorm.DB
}

func NewPostgres() (*Postgres, error) {
	dsn := os.Getenv("DB_URL")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	pg := &Postgres{DB: db}
	err = pg.autoMigrate()
	if err != nil {
		return nil, fmt.Errorf("failed to run auto migration: %w", err)
	}

	return pg, nil
}

func (pg *Postgres) autoMigrate() error {
	err := pg.DB.AutoMigrate(
		&entities.User{},
		&entities.Plan{},
		&entities.Board{},
	)

	if err != nil {
		return err
	}

	return nil
}
