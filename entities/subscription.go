package entities

import "time"

type Subscription struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	EndDate     time.Time `json:"endDate" binding:"required"`
	CreatedDate time.Time `json:"createdDate" binding:"required"`
}
