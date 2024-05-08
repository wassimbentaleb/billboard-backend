package entities

type Board struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	State       string `json:"state"`
	Status      string `json:"status"`
	Plans       []Plan `gorm:"foreignKey:BoardID"`
}
