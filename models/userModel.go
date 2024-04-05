package models

import (
	"github.com/lib/pq"
	"time"
)

type User struct {
	ID        uint `gorm:"primarykey"`
	FirstName string
	LastName  string
	Email     string `gorm:"unique"`
	Password  string
}

type Client struct {
	ID        uint `gorm:"primarykey"`
	Email     string
	FirstName string
	LastName  string
	State     string
}

type Billboard struct {
	ID          uint `gorm:"primarykey"`
	Name        string
	Description string
	State       string
	Status      string
}

type Plans struct {
	BoardId     string
	ID          uint `gorm:"primarykey"`
	Title       string
	StartDate   time.Time
	EndDate     time.Time
	Description string
	ImageUrls   pq.StringArray `gorm:"type:text[]"`
}
