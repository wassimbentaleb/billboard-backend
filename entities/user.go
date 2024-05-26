package entities

type User struct {
	ID                 uint   `json:"id" gorm:"primaryKey"`
	CompanyName        string `json:"company_name" binding:"required"`
	Email              string `json:"email" binding:"required,email" gorm:"unique"`
	Password           string `json:"password" binding:"required,min=8"`
	PhoneNumber        string `json:"phone_number" binding:"required"`
	ActiveSubscription string `json:"active_subscription" binding:"required"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}
