package entities

type Board struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	Ref         string `json:"ref" gorm:"unique"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	Address     string `json:"address"`
	Status      string `json:"status" gorm:"default:not active"`
	Plans       []Plan `gorm:"foreignKey:BoardID"`
	UserID      *uint  `json:"-"`
}
