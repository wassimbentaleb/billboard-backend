package entities

type User struct {
	ID                 uint           `json:"id" gorm:"primaryKey"`
	CompanyName        string         `json:"company_name" binding:"required"`
	Email              string         `json:"email" binding:"required,email" gorm:"unique"`
	Password           string         `json:"-"`
	PhoneNumber        string         `json:"phone_number" binding:"required"`
	Address            string         `json:"address" binding:"required"`
	ActiveSubscription string         `json:"-"`
	IsAdmin            bool           `json:"-" gorm:"default:false"`
	Subscriptions      []Subscription `json:"subscriptions" gorm:"foreignKey:UserID"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type UpdateUserRequest struct {
	CompanyName string `json:"company_name" binding:"required"`
	Email       string `json:"email" binding:"required,email"`
	PhoneNumber string `json:"phone_number" binding:"required"`
	Address     string `json:"address" binding:"required"`
	Password    string `json:"password"`
	NewPassword string `json:"new_password"`
}
