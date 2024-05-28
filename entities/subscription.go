package entities

type Subscription struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	CompanyName string `json:"company_name" binding:"required"`
	EndDate     string `json:"endDate" binding:"required"`
	CreatedDate string `json:"createdDate" binding:"required"`
	Paid        string `json:"paid" binding:"required"`
}
