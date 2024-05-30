package entities

type Subscription struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	EndDate     string `json:"endDate" binding:"required"`
	CreatedDate string `json:"createdDate" binding:"required"`
	Paid        string `json:"paid" binding:"required"`
	UserID      uint   `json:"-"`
}
