package entities

import (
	"time"

	"github.com/lib/pq"
)

type Plan struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	Title       string         `json:"title" binding:"required"`
	StartDate   time.Time      `json:"startDate" binding:"required"`
	EndDate     time.Time      `json:"endDate" binding:"required"`
	Description string         `json:"description"`
	ImageUrls   pq.StringArray `json:"imageUrls" gorm:"type:text[]"`
	BoardID     uint           `json:"boardId"`
}
