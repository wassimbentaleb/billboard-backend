package entities

type Board struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	Ref         string `json:"ref" gorm:"unique"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	Address     string `json:"address" binding:"required"`
	Status      string `json:"status"`
	Plans       []Plan `gorm:"foreignKey:BoardID"`
}
