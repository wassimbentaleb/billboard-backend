package entities

import (
	"time"

	"github.com/lib/pq"
)

type Plan struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	Title       string         `json:"title" binding:"required"`
	StartDate   time.Time      `json:"start_date" binding:"required"`
	EndDate     time.Time      `json:"end_date" binding:"required"`
	Description string         `json:"description"`
	ImageUrls   pq.StringArray `json:"image_urls" gorm:"type:text[]"`
}
