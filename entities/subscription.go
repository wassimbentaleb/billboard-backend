package entities

type Subscription struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	EndDate     string `json:"endDate" binding:"required"`
	CreatedDate string `json:"createdDate" binding:"required"`
	Paid        string `json:"paid" binding:"required"`
	UserID      uint   `json:"company_name"`
	User        *User  `json:"-"`
}

type SubscriptionResponse struct {
	ID          uint   `json:"id"`
	EndDate     string `json:"endDate"`
	CreatedDate string `json:"createdDate"`
	Paid        string `json:"paid"`
	CompanyName string `json:"company_name"`
}
